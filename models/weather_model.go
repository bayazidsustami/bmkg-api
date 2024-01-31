package models

type Weather struct {
	TimeStamp        int64  `json:"timestamp"`
	Source           string `json:source`
	ProductionCenter string `json:production_center`
	Domain           string `json:domain`
	Areas            []Area `json:"areas"`
}

type SingleWeather struct {
	TimeStamp        int64  `json:"timestamp"`
	Source           string `json:source`
	ProductionCenter string `json:production_center`
	Domain           string `json:domain`
	Areas            Area   `json:"area"`
}
