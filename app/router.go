package app

import (
	"github.com/bayazidsustami/bmkg-api/controller"
	"github.com/bayazidsustami/bmkg-api/repository"
	"github.com/bayazidsustami/bmkg-api/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	repository := repository.New()
	service := service.New(repository)
	controller := controller.New(service)

	app.Get("/api/forecast/:provinceId", controller.GetForecastById)
}
