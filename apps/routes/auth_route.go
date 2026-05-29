//- apps/routes/auth_route.go

package routes

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitAuth(app *fiber.App, cfg *configs.Config, db *gorm.DB) {
	handler := handlers.NewAuthHandler(cfg, db)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}
