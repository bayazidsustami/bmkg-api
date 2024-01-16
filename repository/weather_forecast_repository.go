package repository

import "github.com/bayazidsustami/bmkg-api/models"

type WeatherForecastRepository interface {
	GetForecastById(id string) (int, *models.Weather, error)
}
