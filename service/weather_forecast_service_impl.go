package service

type WeatherForecastService interface {
	GetForecastById(id string) error
}
