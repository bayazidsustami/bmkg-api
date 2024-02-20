package controller

import (
	"net/http"
	"strconv"

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
	cityIdInt, _ := strconv.Atoi(id)

	responseErr := checkProvinceIdInvalid(cityIdInt)
	if responseErr != nil {
		return ctx.JSON(responseErr)
	}

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

func (w *WeatherForecastControllerImpl) GetForecastCities(ctx *fiber.Ctx) error {
	statusCode, weather, err := w.Service.GetForecastById("35")

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

func (w *WeatherForecastControllerImpl) GetForecastByCity(ctx *fiber.Ctx) error {
	id := ctx.Params("provinceId")
	cityIdInt, _ := strconv.Atoi(id)

	responseErr := checkProvinceIdInvalid(cityIdInt)
	if responseErr != nil {
		return ctx.JSON(responseErr)
	}

	cityId := ctx.Params("cityId")
	statusCode, weather, err := w.Service.GetForecastByCity(id, cityId)

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

func checkProvinceIdInvalid(id int) *models.ReponseError {
	if id > 35 {
		return &models.ReponseError{
			StatusCode: http.StatusNotFound,
			Message:    "Not Found for Province ID : " + strconv.Itoa(id),
		}
	}
	return nil
}
