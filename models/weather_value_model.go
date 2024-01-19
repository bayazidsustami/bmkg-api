package models

import "github.com/bayazidsustami/bmkg-api/clients"

type WeatherValue struct {
	Id          string          `json:"id"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	Data        []WeatherValues `json:"data"`
}

type WeatherValues struct {
	Humidity    rune   `json:"hour"`
	DateTime    int64  `json:"datetime"`
	Icon        rune   `json:"icon"`
	Description string `json:"description"`
}

func (w *WeatherValues) SetDescription() {
	w.Description = clients.WeatherCode[int(w.Icon)]
}
