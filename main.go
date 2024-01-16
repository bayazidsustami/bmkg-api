package main

import (
	"log"
	"time"

	"github.com/bayazidsustami/bmkg-api/controller"
	"github.com/bayazidsustami/bmkg-api/repository"
	"github.com/bayazidsustami/bmkg-api/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	repository := repository.New()
	service := service.New(repository)
	controller := controller.New(service)

	app.Get("/api/forecast", controller.GetForecastById)

	err := app.Listen(":8000")
	log.Fatal(err)
}
