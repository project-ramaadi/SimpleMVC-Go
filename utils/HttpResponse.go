package utils

type HttpResponse[T any] struct {
	Success    bool
	Message    string
	StatusCode int
	Body       T
}
