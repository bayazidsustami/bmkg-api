package models

type Weather struct {
	TimeStamp int64  `json:"timestamp"`
	Areas     []Area `json:"areas"`
}

type SingleWeather struct {
	TimeStamp int64 `json:"timestamp"`
	Areas     Area  `json:"area"`
}
