package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/typesense/typesense-go/v4/typesense"
	"golang.org/x/crypto/bcrypt"

	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"
	"go-fiber-dummyapi-svc/apps/routes"
	"go-fiber-dummyapi-svc/pkgs/response"
)

// ---------------------------------------------------------------------------
// Test data
// ---------------------------------------------------------------------------

const (
	testUserID    = "1"
	testFirstname = "John"
	testLastname  = "Doe"
	testEmail     = "john@example.com"
	testAvatar    = "https://example.com/avatar.jpg"
	testPassword  = "secret123"
)

var testPasswordHash string

func init() {
	hash, err := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	testPasswordHash = string(hash)
}

// ---------------------------------------------------------------------------
// Test helpers
// ---------------------------------------------------------------------------

func newTestConfig() *configs.Config {
	return &configs.Config{
		Server: configs.ConfigServer{
			AppName: "dummyapi-svc-test",
			Port:    0,
			Prefork: false,
			Debug:   true,
		},
		Auth: configs.ConfigAuth{
			JwtSecret:             "test-secret-key-1234567890",
			JwtAccessTokenExpire:  15,
			JwtRefreshTokenExpire: 1440,
		},
		Typesense: configs.ConfigTypesense{},
	}
}

func generateTestToken(cfg *configs.Config, expireMinutes int) string {
	claims := jwt.MapClaims{
		"id":        testUserID,
		"firstname": testFirstname,
		"lastname":  testLastname,
		"email":     testEmail,
		"avatar":    testAvatar,
		"exp":       time.Now().Add(time.Duration(expireMinutes) * time.Minute).Unix(),
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(cfg.Auth.JwtSecret))
	return signed
}

func newMockTypesenseServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Search endpoint
		if strings.Contains(r.URL.Path, "/documents/search") {
			filterBy := r.URL.Query().Get("filter_by")

			// If filter_by contains our test email, return a match
			if strings.Contains(filterBy, testEmail) {
				doc := map[string]interface{}{
					"id":            testUserID,
					"firstname":     testFirstname,
					"lastname":      testLastname,
					"email":         testEmail,
					"avatar":        testAvatar,
					"password_hash": testPasswordHash,
				}
				found := 1
				hits := []map[string]interface{}{
					{"document": doc},
				}
				json.NewEncoder(w).Encode(map[string]interface{}{
					"found":  found,
					"hits":   hits,
					"out_of": found,
				})
				return
			}

			// If filter_by exists but not our email, return no results
			if filterBy != "" {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"found":  0,
					"out_of": 0,
				})
				return
			}

			// General list search (no filter_by)
			perPage := r.URL.Query().Get("per_page")
			if perPage == "" {
				perPage = "20"
			}

			users := []map[string]interface{}{
				{
					"id":        "1",
					"firstname": "John",
					"lastname":  "Doe",
					"email":     "john@example.com",
					"avatar":    testAvatar,
					"gender":    "M",
					"phone":     "1234567890",
				},
				{
					"id":        "2",
					"firstname": "Jane",
					"lastname":  "Smith",
					"email":     "jane@example.com",
					"avatar":    testAvatar,
					"gender":    "F",
					"phone":     "0987654321",
				},
			}

			var hits []map[string]interface{}
			for _, u := range users {
				hits = append(hits, map[string]interface{}{
					"document": u,
				})
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"found":  len(users),
				"out_of": len(users),
				"page":   1,
				"hits":   hits,
			})
			return
		}

		// Document retrieve endpoint: /collections/dummy_users/documents/{id}
		parts := strings.Split(r.URL.Path, "/")
		docID := parts[len(parts)-1]

		switch docID {
		case "1":
			json.NewEncoder(w).Encode(map[string]interface{}{
				"id":            "1",
				"firstname":     "John",
				"lastname":      "Doe",
				"email":         "john@example.com",
				"avatar":        testAvatar,
				"password":      testPassword,
				"password_hash": testPasswordHash,
				"gender":        "M",
				"phone":         "1234567890",
			})
		default:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Not Found",
			})
		}
	}))
}

func newTestApp(cfg *configs.Config, ts *typesense.Client) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	routes.RouteAuth(app, cfg, ts)
	routes.RouteUser(app, cfg, ts)
	app.Use(func(c *fiber.Ctx) error {
		return handlers.RespError(c, 404, "Route Not Found", nil)
	})
	return app
}

type apiResponse struct {
	Meta    response.ResponseMeta `json:"meta"`
	Message string                `json:"message"`
	Record  interface{}           `json:"record"`
	Records interface{}           `json:"records"`
	Errors  interface{}           `json:"errors"`
}

func decodeResponse(t *testing.T, resp *http.Response) apiResponse {
	t.Helper()
	var result apiResponse
	err := json.NewDecoder(resp.Body).Decode(&result)
	require.NoError(t, err)
	return result
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

func TestAuthLogin(t *testing.T) {
	t.Parallel()

	tsServer := newMockTypesenseServer()
	defer tsServer.Close()

	cfg := newTestConfig()
	tsClient := typesense.NewClient(
		typesense.WithServer(tsServer.URL),
		typesense.WithAPIKey("test-key"),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	app := newTestApp(cfg, tsClient)

	t.Run("success", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, testEmail, testPassword)
		req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, "", result.Message)

		record, ok := result.Record.(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, testUserID, record["id"])
		assert.Equal(t, testFirstname, record["firstname"])
		assert.Equal(t, testEmail, record["email"])
		assert.NotEmpty(t, record["accessToken"])
		assert.NotEmpty(t, record["refreshToken"])
	})

	t.Run("user not found", func(t *testing.T) {
		body := `{"email":"unknown@example.com","password":"secret"}`
		req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, "User tidak ditemukan", result.Message)
	})

	t.Run("wrong password", func(t *testing.T) {
		body := fmt.Sprintf(`{"email":"%s","password":"wrongpass"}`, testEmail)
		req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, "Password salah", result.Message)
	})
}

func TestAuthRefreshToken(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig()
	configs.Cfg = cfg // RefreshToken handler uses global configs.Cfg

	tsServer := newMockTypesenseServer()
	defer tsServer.Close()

	tsClient := typesense.NewClient(
		typesense.WithServer(tsServer.URL),
		typesense.WithAPIKey("test-key"),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	app := newTestApp(cfg, tsClient)

	t.Run("success", func(t *testing.T) {
		refreshToken := generateTestToken(cfg, cfg.Auth.JwtRefreshTokenExpire)
		body := fmt.Sprintf(`{"refreshToken":"%s"}`, refreshToken)
		req := httptest.NewRequest(http.MethodPost, "/auth/refresh-token", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, "", result.Message)

		record, ok := result.Record.(map[string]interface{})
		require.True(t, ok)
		assert.NotEmpty(t, record["accessToken"])
		assert.NotEmpty(t, record["refreshToken"])
	})

	t.Run("invalid token", func(t *testing.T) {
		body := `{"refreshToken":"invalid-token"}`
		req := httptest.NewRequest(http.MethodPost, "/auth/refresh-token", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

func TestAuthMe(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig()
	tsServer := newMockTypesenseServer()
	defer tsServer.Close()

	tsClient := typesense.NewClient(
		typesense.WithServer(tsServer.URL),
		typesense.WithAPIKey("test-key"),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	app := newTestApp(cfg, tsClient)

	t.Run("success", func(t *testing.T) {
		accessToken := generateTestToken(cfg, cfg.Auth.JwtAccessTokenExpire)
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		req.Header.Set("Authorization", "Bearer "+accessToken)

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)

		record, ok := result.Record.(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, testUserID, record["id"])
		assert.Equal(t, testFirstname, record["firstname"])
		assert.Equal(t, testEmail, record["email"])
	})

	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, "Missing token", result.Message)
	})

	t.Run("expired token", func(t *testing.T) {
		expiredToken := generateTestToken(cfg, -1) // expired
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		req.Header.Set("Authorization", "Bearer "+expiredToken)

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("invalid token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/me", nil)
		req.Header.Set("Authorization", "Bearer some-invalid-token")

		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

func TestUserList(t *testing.T) {
	t.Parallel()

	tsServer := newMockTypesenseServer()
	defer tsServer.Close()

	cfg := newTestConfig()
	configs.Cfg = cfg

	tsClient := typesense.NewClient(
		typesense.WithServer(tsServer.URL),
		typesense.WithAPIKey("test-key"),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	app := newTestApp(cfg, tsClient)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	resp, err := app.Test(req, 5000)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result apiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	require.NoError(t, err)

	records, ok := result.Records.([]interface{})
	require.True(t, ok)
	assert.GreaterOrEqual(t, len(records), 1)

	first := records[0].(map[string]interface{})
	assert.Equal(t, "John", first["firstname"])
}

func TestUserDetail(t *testing.T) {
	t.Parallel()

	tsServer := newMockTypesenseServer()
	defer tsServer.Close()

	cfg := newTestConfig()
	configs.Cfg = cfg

	tsClient := typesense.NewClient(
		typesense.WithServer(tsServer.URL),
		typesense.WithAPIKey("test-key"),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	app := newTestApp(cfg, tsClient)

	t.Run("found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)

		record, ok := result.Record.(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, "1", record["id"])
		assert.Equal(t, "John", record["firstname"])
	})

	t.Run("not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
		resp, err := app.Test(req, 5000)
		require.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var result apiResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		assert.Equal(t, "Data not found", result.Message)
	})
}

func TestRouteNotFound(t *testing.T) {
	t.Parallel()

	cfg := newTestConfig()
	tsServer := newMockTypesenseServer()
	defer tsServer.Close()

	tsClient := typesense.NewClient(
		typesense.WithServer(tsServer.URL),
		typesense.WithAPIKey("test-key"),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	app := newTestApp(cfg, tsClient)

	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	resp, err := app.Test(req, 5000)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var result apiResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	require.NoError(t, err)
	assert.Equal(t, "Route Not Found", result.Message)
}
