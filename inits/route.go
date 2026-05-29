//- inits/route.go

package inits

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, cfg *configs.Config, db *gorm.DB, ts *typesense.Client) {
	routes.InitAuth(app, cfg, db)
	routes.RouteUser(app, cfg, ts)
}
