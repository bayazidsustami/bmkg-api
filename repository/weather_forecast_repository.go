package repository

type WeatherForecastRepository interface {
	GetForecastById(id string) (int, string, error)
}
