package repository

import (
	"net/http"

	"github.com/bayazidsustami/bmkg-api/clients"
)

type WeatherForecastRepositoryImpl struct {
}

func New() WeatherForecastRepository {
	return &WeatherForecastRepositoryImpl{}
}

func (w *WeatherForecastRepositoryImpl) GetForecastById(id string) (int, string, error) {
	client := clients.GetClient()

	agent := client.Get(clients.GetPath(id))
	status, response, errs := agent.String()
	if errs != nil {
		return status, "", errs[0]
	}

	return http.StatusOK, response, nil
}
