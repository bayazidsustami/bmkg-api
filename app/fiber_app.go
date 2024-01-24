package app

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StartFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	})

	RegisterRoute(app)

	err := app.Listen(":8000")
	log.Fatal(err)
}
