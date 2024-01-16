package models

type Humidity struct {
	Id          string          `json:"id"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	Data        []HumidityValue `json:"data"`
}

type HumidityValue struct {
	Humidity rune   `json:"h"`
	DateTime int64  `json:"datetime"`
	Value    rune   `json:"value"`
	Unit     string `json:"unit"`
}
