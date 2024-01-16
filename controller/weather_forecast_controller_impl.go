package controller

import (
	"github.com/bayazidsustami/bmkg-api/service"
	"github.com/gofiber/fiber/v2"
)

type WeatherForecastControllerImpl struct {
	Service service.WeatherForecastService
}

func New(service service.WeatherForecastService) WeatherForecastController {
	return &WeatherForecastControllerImpl{
		Service: service,
	}
}

func (w *WeatherForecastControllerImpl) GetForecastById(ctx *fiber.Ctx) error {
	return w.Service.GetForecastById("1")
}
