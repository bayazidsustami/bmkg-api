package models

type WindSpeed struct {
	Id          string           `json:"id"`
	Description string           `json:"description"`
	Type        string           `json:"type"`
	Data        []WindSpeedValue `json:"data"`
}

type WindSpeedValue struct {
	Hour     rune           `json:"hour"`
	DateTime int64          `json:"datetime"`
	Values   map[string]any `json:"values"`
}
