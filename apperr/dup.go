package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

// Duplicate wraps Duplicate error
type Duplicate interface {
	Duplicate() error
}

type duplicate struct {
	What  string
	Where string
}

// StatusCode returns http status code
func (d duplicate) StatusCode() int {
	return http.StatusBadRequest
}

func (d duplicate) Error() string {
	return d.String()
}

func (d duplicate) Name() string {
	return DuplicateN
}

func (d duplicate) String() string {
	return fmt.Sprintf("%s#where=%s#what=%s", d.Name(), d.Where, d.What)
}

func (d duplicate) Duplicate() error {
	return errors.New(d.String())
}

// NewDuplicate return Duplicate error
func NewDuplicate(where, what string) error {
	return duplicate{
		What:  what,
		Where: where,
	}
}
