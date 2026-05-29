//- apps/routes/user_route.go

package routes

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

func RouteUser(api fiber.Router, cfg *configs.Config, ts *typesense.Client) {
	h := handlers.NewUserHandler(ts)
	ep := "/users"

	api.Get(ep, h.List)
	api.Get(ep+"/:id", h.Detail)
}
