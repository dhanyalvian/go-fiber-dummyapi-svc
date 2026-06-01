//- apps/middlewares/auth_middleware.go

package middlewares

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"
	"go-fiber-dummyapi-svc/pkgs/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Protected(cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Ambil token dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return handlers.RespError(c, 401, "Missing token", nil)
		}

		// 2. Format: "Bearer <token>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 3. Parsing dan Validasi Token
		tokenData, err := utils.DecodeToken(cfg, tokenString)
		if err != nil {
			return handlers.RespError(c, 401, err.Error(), err)
		}

		// 4. Simpan data user ke dalam context agar bisa dipakai di handler
		c.Locals("user_id", tokenData.ID)
		c.Locals("user_token", tokenData)

		return c.Next()
	}
}
