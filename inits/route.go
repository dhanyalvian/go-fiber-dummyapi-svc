//- inits/route.go

package inits

import (
	"go-fiber-dummy-svc/apps/configs"
	"go-fiber-dummy-svc/apps/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, cfg *configs.Config, db *gorm.DB) {
	routes.InitAuth(app, cfg, db)
	routes.RouteUser(app, db)
}
