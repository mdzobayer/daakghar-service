package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

// Valider wraps the validation error
type Valider interface {
	Valid() error
}

type validation struct {
	What  string
	Cause string
	Hint  string
}

// StatusCode returns http status code
func (v validation) StatusCode() int {
	return http.StatusBadRequest
}

func (v validation) Error() string {
	return v.String()
}

func (v validation) Name() string {
	return ValidationN
}

func (v validation) String() string {
	return fmt.Sprintf("%s#what=%s,cause=%s,hint=%s", v.Name(), v.What, v.Cause, v.Hint)
}

func (v validation) Valid() error {
	return errors.New(v.String())
}

// NewValidation return validation error
func NewValidation(what, cause, hint string) error {
	return validation{
		What:  what,
		Cause: cause,
		Hint:  hint,
	}
}
