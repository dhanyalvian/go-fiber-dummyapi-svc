//- inits/database.go

package inits

import (
	"fmt"
	"log"
	"time"

	"go-fiber-dummyapi-svc/apps/configs"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb(cfg *configs.Config) *gorm.DB {
	var err error

	cfgDb := cfg.Database
	cfgDbType := cfgDb.Type

	switch cfgDbType {
	case "mysql":
		Db, err = initMysql(cfgDb)
	case "pgsql":
		Db, err = initPostgresql(cfgDb)
	default:
		log.Fatal("Database not supported...", err)
	}

	if err != nil {
		log.Fatal("Failed to Connect to database...", err)
	}

	if cfgDb.Debug {
		Db = Db.Debug()
		println("🔧 GORM Debug Mode: Enabled")
	}

	return optimizeConnection(Db)
}

func initMysql(cfgDb configs.ConfigDatabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfgDb.Username,
		cfgDb.Password,
		cfgDb.Hostname,
		cfgDb.Port,
		cfgDb.DbName,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func initPostgresql(cfgDb configs.ConfigDatabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfgDb.Hostname,
		cfgDb.Username,
		cfgDb.Password,
		cfgDb.DbName,
		cfgDb.Port,
		cfgDb.SslMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func optimizeConnection(db *gorm.DB) *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get DB object:", err)
	}

	// SetMaxIdleConns menetapkan jumlah maksimum koneksi dalam pool koneksi yang tidak digunakan.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns menetapkan jumlah maksimum koneksi terbuka ke database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime menetapkan jumlah waktu maksimum koneksi dapat digunakan kembali.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetConnMaxIdleTime sets the maximum amount of time a connection can be idle before being closed.
	sqlDB.SetConnMaxIdleTime(time.Hour)

	// Use the 'db' object throughout your application
	fmt.Println("Database connected and connection pool configured!")

	return db
}
