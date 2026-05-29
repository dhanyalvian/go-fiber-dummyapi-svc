//- apps/configs/config.go

package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ConfigServer    `mapstructure:",squash"`
	Auth      ConfigAuth      `mapstructure:",squash"`
	Database  ConfigDatabase  `mapstructure:",squash"`
	Typesense ConfigTypesense `mapstructure:",squash"`
}

type ConfigServer struct {
	AppName  string `mapstructure:"APP_NAME"`
	Port     int    `mapstructure:"APP_PORT"`
	Prefork  bool   `mapstructure:"APP_PREFORK"`
	Debug    bool   `mapstructure:"APP_DEBUG"`
	ImageURL string `mapstructure:"APP_IMAGE_URL"`
}

type ConfigDatabase struct {
	Type     string `mapstructure:"DB_TYPE"`
	Hostname string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	DbName   string `mapstructure:"DB_NAME"`
	SslMode  string `mapstructure:"DB_SSLMODE"`
	TimeZone string `mapstructure:"DB_TIMEZONE"`
	Debug    bool   `mapstructure:"DB_DEBUG"`
}

type ConfigTypesense struct {
	Active   bool   `mapstructure:"TS_ACTIVE"`
	Hostname string `mapstructure:"TS_HOST"`
	Port     int    `mapstructure:"TS_PORT"`
	ApiKey   string `mapstructure:"TS_KEY"`
}

type ConfigAuth struct {
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

var Cfg *Config

func Get() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	viper.AutomaticEnv()

	Cfg = &Config{}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return Cfg
}
