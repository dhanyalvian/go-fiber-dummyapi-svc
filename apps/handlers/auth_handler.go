//- apps/handlers/auth_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/apps/models"
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Cfg *configs.Config
	TS  *typesense.Client
}

func NewAuthHandler(cfg *configs.Config, ts *typesense.Client) *AuthHandler {
	return &AuthHandler{Cfg: cfg, TS: ts}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	docs, err := models.Login(c, h.TS, data["email"])
	if *docs.Found == 0 {
		return RespError(c, 400, "User tidak ditemukan", err)
	}

	user := GetDocFirst[entities.RespDetailUser](docs)

	// Bandingkan password
	if !CheckPasswordHash(data["password"], user.PasswordHash) {
		return RespError(c, 400, "Password salah", err)
	}

	// Generate JWT Token
	tokenData := utils.TokenData{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	// Generate kedua token secara paralel
	type tokenResult struct {
		token string
		err   error
	}

	accessCh := make(chan tokenResult, 1)
	refreshCh := make(chan tokenResult, 1)

	go func() {
		token, err := utils.GenerateToken(h.Cfg, tokenData, h.Cfg.Auth.JwtAccessTokenExpire)
		accessCh <- tokenResult{token, err}
	}()
	go func() {
		token, err := utils.GenerateToken(h.Cfg, tokenData, h.Cfg.Auth.JwtRefreshTokenExpire)
		refreshCh <- tokenResult{token, err}
	}()

	accessResult := <-accessCh
	refreshResult := <-refreshCh

	if accessResult.err != nil {
		return RespError(c, 500, accessResult.err.Error(), accessResult.err)
	}
	if refreshResult.err != nil {
		return RespError(c, 500, refreshResult.err.Error(), refreshResult.err)
	}

	result := entities.RespAuthLogin{
		BaseID: entities.BaseID{
			ID: user.ID,
		},
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		Avatar:       user.Avatar,
		AccessToken:  accessResult.token,
		RefreshToken: refreshResult.token,
	}

	return RespSucess(c, "", result, nil, nil)
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return RespError(c, 401, err.Error(), err)
	}

	tokenDecode, err := utils.DecodeToken(configs.Cfg, data["refreshToken"])
	if err != nil {
		return RespError(c, 401, err.Error(), err)
	}

	tokenData := utils.TokenData{
		ID:        tokenDecode.ID,
		Firstname: tokenDecode.Firstname,
		Lastname:  tokenDecode.Lastname,
		Email:     tokenDecode.Email,
		Avatar:    tokenDecode.Avatar,
	}

	accessTokenExpire := h.Cfg.Auth.JwtAccessTokenExpire
	accessToken, err := utils.GenerateToken(h.Cfg, tokenData, accessTokenExpire)
	if err != nil {
		return RespError(c, 500, err.Error(), err)
	}

	refreshTokenExpire := h.Cfg.Auth.JwtRefreshTokenExpire
	refreshToken, err := utils.GenerateToken(h.Cfg, tokenData, refreshTokenExpire)
	if err != nil {
		return RespError(c, 500, err.Error(), err)
	}

	result := fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	return RespSucess(c, "", result, nil, nil)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
