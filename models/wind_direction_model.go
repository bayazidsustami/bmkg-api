package models

type WindDirection struct {
	Id          string               `json:"id"`
	Description string               `json:"description"`
	Type        string               `json:"type"`
	Data        []WindDirectionValue `json:"data"`
}

type WindDirectionValue struct {
	Hour     rune           `json:"hour"`
	DateTime int64          `json:"datetime"`
	Values   map[string]any `json:"values"`
}
