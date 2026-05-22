//- cmd/api/main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"go-fiber-dummy-svc/apps/configs"
	"go-fiber-dummy-svc/apps/databases"
	"go-fiber-dummy-svc/inits"

	"github.com/gofiber/fiber/v2"
	"github.com/typesense/typesense-go/v4/typesense"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.Get()

	app := fiber.New(fiber.Config{
		AppName:     cfg.Server.AppName,
		Prefork:     cfg.Server.Prefork,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	db := inits.InitDb(cfg)
	ts := inits.InitTs(cfg)

	handleArgs(db)

	runServer(app, cfg, db, ts)
}

func handleArgs(db *gorm.DB) {
	// Definisikan flag --migrate
	migrateFlag := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	if *migrateFlag {
		databases.Migration(db)
		os.Exit(0) // Berhenti setelah migrasi selesai
	}
}

func runServer(app *fiber.App, cfg *configs.Config, db *gorm.DB, ts *typesense.Client) {
	inits.InitApp(app)
	inits.InitLogger(app)
	inits.InitRouter(app, cfg, db, ts)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
