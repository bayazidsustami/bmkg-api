package models

type Weather struct {
	Issue Issue  `json:"issue"`
	Areas []Area `json:"areas"`
}

type Issue struct {
	TimeStamp int64 `json:"timestamp"`
	Year      rune  `json:"year"`
	Month     rune  `json:"month"`
	Day       rune  `json:"day"`
	Hour      rune  `json:"Hour"`
	Minute    rune  `json:"Minute"`
	Second    rune  `json:"Second"`
}

type Area struct {
	Id          rune    `json:"id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Coordinate  string  `json:"coordinate"`
	Type        string  `json:"type"`
	Region      string  `json:"region"`
	Level       rune    `json:"level"`
	Description string  `json:"description"`
	Domain      string  `json:"string"`
	Tags        string  `json:"tags"`
}
