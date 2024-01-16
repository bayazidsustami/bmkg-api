package repository

type WeatherForecastRepository interface {
	GetForecastById(id string) error
}
