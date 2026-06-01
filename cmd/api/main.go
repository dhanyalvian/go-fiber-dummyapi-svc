//- cmd/api/main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/databases"
	"go-fiber-dummyapi-svc/apps/databases/seeders"
	"go-fiber-dummyapi-svc/inits"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
)

func main() {
	cfg := configs.Get()

	app := fiber.New(fiber.Config{
		AppName:     cfg.Server.AppName,
		Prefork:     cfg.Server.Prefork,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	ts := inits.InitTs(cfg)

	handleArgs(ts)
	runServer(app, cfg, ts)
}

func handleArgs(ts *typesense.Client) {
	migrateFlag := flag.Bool("migrate", false, "Run Typesense migrations and seed data")
	flag.Parse()

	if *migrateFlag {
		databases.MigrateTypesense(ts)
		seeders.SeedAll(ts)
		os.Exit(0)
	}
}

func runServer(app *fiber.App, cfg *configs.Config, ts *typesense.Client) {
	inits.InitApp(app)
	inits.InitLogger(app)
	inits.InitRouter(app, cfg, ts)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
