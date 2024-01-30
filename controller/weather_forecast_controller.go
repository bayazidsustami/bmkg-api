package controller

import "github.com/gofiber/fiber/v2"

type WeatherForecastController interface {
	GetForecastById(ctx *fiber.Ctx) error
	GetForecastCities(ctx *fiber.Ctx) error
	GetForecastByCity(ctx *fiber.Ctx) error
}
