package controller

import "github.com/gofiber/fiber/v2"

type WeatherForecastControllerImpl struct {
}

func New() WeatherForecastController {
	return &WeatherForecastControllerImpl{}
}

func (w *WeatherForecastControllerImpl) GetForecastById(ctx *fiber.Ctx) error {
	return nil
}
