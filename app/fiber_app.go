package app

import (
	"log"

	"github.com/bayazidsustami/bmkg-api/utils"
	"github.com/gofiber/fiber/v2"
)

func StartFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  utils.IdleTimeout,
		WriteTimeout: utils.WriteTimeout,
		ReadTimeout:  utils.ReadTimeout,
	})

	RegisterRoute(app)

	err := app.Listen(":8000")
	log.Fatal(err)
}
