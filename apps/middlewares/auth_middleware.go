//- apps/middlewares/auth_middleware.go

package middlewares

import (
	"strings"

	"go-fiber-dummy-svc/apps/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected(cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Ambil token dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing token"})
		}

		// 2. Format: "Bearer <token>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 3. Parsing dan Validasi Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Auth.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
		}

		// 4. Simpan data user ke dalam context agar bisa dipakai di handler
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["id"])

		return c.Next()
	}
}
