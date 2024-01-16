package models

type Weather struct {
	TimeStamp int64  `json:"timestamp"`
	Areas     []Area `json:"areas"`
}

type Area struct {
	Id          rune     `json:"id"`
	Name        string   `json:"name"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Coordinate  string   `json:"coordinate"`
	Type        string   `json:"type"`
	Region      string   `json:"region"`
	Level       rune     `json:"level"`
	Description string   `json:"description"`
	Domain      string   `json:"domain"`
	Tags        string   `json:"tags"`
	Humidy      Humidity `json:"humidity"`
}

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
