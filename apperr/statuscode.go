package apperr

// StatusCoder wraps http status code
type StatusCoder interface {
	StatusCode() int
}
