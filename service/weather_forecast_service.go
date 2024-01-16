package service

import (
	"net/http"

	"github.com/bayazidsustami/bmkg-api/models"
	"github.com/bayazidsustami/bmkg-api/repository"
	"github.com/bayazidsustami/bmkg-api/utils"
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
	code, response, err := service.Repository.GetForecastById(id)
	if err != nil {
		return code, nil, err
	}
	weather, err := utils.ParseAllElement(response)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return code, weather, nil
}
