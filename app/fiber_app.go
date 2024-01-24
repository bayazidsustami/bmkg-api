package app

import (
	"log"

	"github.com/bayazidsustami/bmkg-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  utils.IdleTimeout,
		WriteTimeout: utils.WriteTimeout,
		ReadTimeout:  utils.ReadTimeout,
	})

	RegisterRoute(app)

	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("env")
	config.AddConfigPath("config/")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(config.GetString("APP_HOST") + ":" + config.GetString("APP_PORT"))
	log.Fatal(err)
}
