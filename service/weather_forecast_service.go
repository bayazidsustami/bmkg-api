package service

import (
	"github.com/bayazidsustami/bmkg-api/models"
	"github.com/bayazidsustami/bmkg-api/repository"
)

type WeatherForecastServiceImpl struct {
	Repository repository.WeatherForecastRepository
}

func New(repository repository.WeatherForecastRepository) WeatherForecastService {
	return &WeatherForecastServiceImpl{
		Repository: repository,
	}
}

func (service *WeatherForecastServiceImpl) GetForecastById(id string) (int, *models.Weather, error) {
	return service.Repository.GetForecastById(id)
}
