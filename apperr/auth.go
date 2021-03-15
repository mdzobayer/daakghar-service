package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

// Auther wraps authentication error
type Auther interface {
	Auth() error
}

type authentication struct {
	Actin string
}

// StatusCode returns http status code
func (a authentication) StatusCode() int {
	return http.StatusUnauthorized
}

func (a authentication) Error() string {
	return a.String()
}

func (a authentication) Name() string {
	return AuthenticationN
}

func (a authentication) String() string {
	return fmt.Sprintf("%s#action=%s", a.Name(), a.Actin)
}

func (a authentication) Auth() error {
	return errors.New(a.String())
}

// NewAuthentication return validation error
func NewAuthentication(action string) error {
	return authentication{
		Actin: action,
	}
}
