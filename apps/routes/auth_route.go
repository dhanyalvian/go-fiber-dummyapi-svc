//- apps/routes/auth_route.go

package routes

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

func RouteAuth(app *fiber.App, cfg *configs.Config, ts *typesense.Client) {
	handler := handlers.NewAuthHandler(cfg, ts)

	app.Post("/login", handler.Login)
	app.Post("/refresh-token", handler.RefreshToken)
}
