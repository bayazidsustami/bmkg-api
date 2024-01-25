package app

import (
	"log"

	"github.com/bayazidsustami/bmkg-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  config.IdleTimeout,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
	})

	RegisterRoute(app)

	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(config.GetString("APP_HOST") + ":" + config.GetString("APP_PORT"))
	log.Fatal(err)
}
