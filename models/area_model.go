package models

type Area struct {
	Id          rune      `json:"id"`
	Name        string    `json:"name"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Coordinate  string    `json:"coordinate"`
	Type        string    `json:"type"`
	Region      string    `json:"region"`
	Level       rune      `json:"level"`
	Description string    `json:"description"`
	Domain      string    `json:"domain"`
	Tags        string    `json:"tags"`
	Parameter   Parameter `json:"parameter"`
}

type Parameter struct {
	Humidity Humidity `json:"humidity"`
}
