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

	weatherGroup := app.Group("/api/weather")

	weatherGroup.Get("/provinces", controller.GetForecastCities)
	weatherGroup.Get("/provinces/:provinceId", controller.GetForecastById)
	weatherGroup.Get("/provinces/:provinceId/cities", controller.GetForecastById)
	weatherGroup.Get("/provinces/:provinceId/cities/:cityId", controller.GetForecastByCity)
}
