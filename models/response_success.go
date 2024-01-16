package models

type ReponseSuccessWithData struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ReponseSuccessWithoutData struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
