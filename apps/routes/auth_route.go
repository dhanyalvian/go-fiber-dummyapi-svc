//- apps/routes/auth_route.go

package routes

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"
	"go-fiber-dummyapi-svc/apps/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

func RouteAuth(app *fiber.App, cfg *configs.Config, ts *typesense.Client) {
	h := handlers.NewAuthHandler(cfg, ts)
	ep := "/auth"

	app.Post(ep+"/login", h.Login)
	app.Post(ep+"/refresh-token", h.RefreshToken)

	api := app.Group(ep, middlewares.Protected(cfg))
	api.Get("/me", h.Me)

}
