package repository

import (
	"net/http"

	"github.com/bayazidsustami/bmkg-api/clients"
	"github.com/bayazidsustami/bmkg-api/models"
	"github.com/bayazidsustami/bmkg-api/utils"
)

type WeatherForecastRepositoryImpl struct {
}

func New() WeatherForecastRepository {
	return &WeatherForecastRepositoryImpl{}
}

func (w *WeatherForecastRepositoryImpl) GetForecastById(id string) (int, *models.Weather, error) {
	client := clients.GetClient()

	agent := client.Get("https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-Aceh.xml")
	status, response, errs := agent.String()
	if errs != nil {
		return status, nil, errs[0]
	}

	weather, err := utils.ParseAllElement(response)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, weather, nil
}
