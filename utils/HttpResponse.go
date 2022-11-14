package utils

type HttpResponse[T any] struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Body       T      `json:"body"`
}
