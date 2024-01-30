package service

import "github.com/bayazidsustami/bmkg-api/models"

type WeatherForecastService interface {
	GetForecastById(id string) (int, *models.Weather, error)
	GetForecastByCity(id string, cityId string) (int, *models.SingleWeather, error)
}
