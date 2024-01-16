package controller

import "github.com/gofiber/fiber/v2"

type WeatherForecastController interface {
	GetForecastById(ctx *fiber.Ctx) error
}
