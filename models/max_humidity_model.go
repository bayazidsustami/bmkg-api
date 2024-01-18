package models

type MaxHumidity struct {
	Id          string             `json:"id"`
	Description string             `json:"description"`
	Type        string             `json:"type"`
	Data        []MaxHumidityValue `json:"data"`
}

type MaxHumidityValue struct {
	DateTime int64  `json:"datetime"`
	Value    rune   `json:"value"`
	Unit     string `json:"unit"`
}
