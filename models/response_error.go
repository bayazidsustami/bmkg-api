package models

type ReponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
