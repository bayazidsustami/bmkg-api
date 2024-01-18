package models

type Temperature struct {
	Id          string             `json:"id"`
	Description string             `json:"description"`
	Type        string             `json:"type"`
	Data        []TemperatureValue `json:"data"`
}

type TemperatureValue struct {
	Humidity rune   `json:"hour"`
	DateTime int64  `json:"datetime"`
	Value    rune   `json:"value"`
	Unit     string `json:"unit"`
}
