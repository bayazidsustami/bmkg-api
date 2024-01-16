package controller

import (
	"github.com/bayazidsustami/bmkg-api/models"
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
	id := ctx.Params("provinceId")
	statusCode, weather, err := w.Service.GetForecastById(id)

	if err != nil {
		return ctx.JSON(models.ReponseError{
			StatusCode: statusCode,
			Message:    err.Error(),
		})
	}

	return ctx.JSON(models.ReponseSuccessWithData{
		StatusCode: statusCode,
		Message:    "Success",
		Data:       weather,
	})
}
