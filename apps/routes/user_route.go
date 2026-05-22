//- apps/routes/user_route.go

package routes

import (
	"go-fiber-dummy-svc/apps/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteUser(api fiber.Router, db *gorm.DB) {
	h := handlers.NewUserHandler(db)
	ep := "/users"

	api.Get(ep, h.List)
	api.Get(ep+"/:id", h.Detail)
}
