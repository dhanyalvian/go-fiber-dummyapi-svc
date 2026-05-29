//- inits/app.go

package inits

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitApp(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(csrf.New(csrf.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/register" || c.Path() == "/login" || c.Path() == "/refresh-token"
		},
	}))
	app.Use(helmet.New())
	app.Use(requestid.New())
}
