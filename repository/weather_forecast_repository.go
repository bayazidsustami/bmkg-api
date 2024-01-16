package repository

import (
	"fmt"

	"github.com/bayazidsustami/bmkg-api/clients"
)

type WeatherForecastRepositoryImpl struct {
}

func New() WeatherForecastRepository {
	return &WeatherForecastRepositoryImpl{}
}

func (w *WeatherForecastRepositoryImpl) GetForecastById(id string) error {
	client := clients.GetClient()

	agent := client.Get("https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-Aceh.xml")
	status, response, errs := agent.String()
	if errs != nil {
		return errs[len(errs)-1]
	}

	fmt.Println(status)
	fmt.Println(response)

	return nil
}
