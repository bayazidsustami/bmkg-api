package models

type MinMaxHumidity struct {
	Id          string                `json:"id"`
	Description string                `json:"description"`
	Type        string                `json:"type"`
	Data        []MinMaxHumidityValue `json:"data"`
}

type MinMaxHumidityValue struct {
	DateTime int64  `json:"datetime"`
	Value    rune   `json:"value"`
	Unit     string `json:"unit"`
}
