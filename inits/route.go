//- inits/route.go

package inits

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"
	"go-fiber-dummyapi-svc/apps/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

func InitRouter(app *fiber.App, cfg *configs.Config, ts *typesense.Client) {
	routes.RouteAuth(app, cfg, ts)
	routes.RoutePost(app, cfg, ts)
	routes.RouteProduct(app, cfg, ts)
	routes.RouteRecipe(app, cfg, ts)
	routes.RouteUser(app, cfg, ts)

	RouteNotFound(app)
}

func RouteNotFound(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return handlers.RespError(c, 404, "Route Not Found", nil)
	})
}
