package models

type MinMaxTemperature struct {
	Id          string                   `json:"id"`
	Description string                   `json:"description"`
	Type        string                   `json:"type"`
	Data        []MinMaxTemperatureValue `json:"data"`
}

type MinMaxTemperatureValue struct {
	DateTime int64  `json:"datetime"`
	Value    rune   `json:"value"`
	Unit     string `json:"unit"`
}
