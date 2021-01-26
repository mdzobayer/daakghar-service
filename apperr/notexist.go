package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

// NotFound wraps NotFound error
type NotFound interface {
	NotFound() error
}

type notfound struct {
	What  string
	Value string
}

// StatusCode returns http status code
func (n notfound) StatusCode() int {
	return http.StatusNotFound
}

func (n notfound) Error() string {
	return n.String()
}

func (n notfound) Name() string {
	return NotFoundN
}

func (n notfound) String() string {
	return fmt.Sprintf("%s#what=%s#value=%s", n.Name(), n.What, n.Value)
}

func (n notfound) NotFoundN() error {
	return errors.New(n.String())
}

// NewNotFound return NotFound error
func NewNotFound(what, value string) error {
	return notfound{
		What:  what,
		Value: value,
	}
}
