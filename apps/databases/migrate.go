//- apps/databases/migrate.go

package databases

import (
	"fmt"
	"log"

	"go-fiber-dummy-svc/apps/entities"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	fmt.Println("Running migrations...")

	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	fmt.Println("Migration completed successfully!")
}
