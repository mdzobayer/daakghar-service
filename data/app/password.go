package app

import (
	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/hash"
)

// Password represent user password
type Password string

// Hash returns hash from password
func (p Password) Hash() (string, error) {
	return hash.GenerateHashFromPassword(string(p))
}

// Valid valids user password
func (p Password) Valid() error {
	ep := validPass{
		password: p,
	}

	ep.minLenValid()
	ep.maxLenValid()

	return ep.Err()
}

// validPass handles password validation
type validPass struct {
	password Password
	err      error
}

// Err returns validPass errors
func (e *validPass) Err() error {
	return e.err
}

// minLenValid checks the minmul length validity of a password
func (e *validPass) minLenValid() {
	if e.err != nil {
		return
	}

	if len(e.password) >= 8 {
		return
	}

	e.err = apperr.NewValidation("password", "length", "minimum is 8")
}

// maxLenValid checks the minmul length validity of a password
func (e *validPass) maxLenValid() {
	if e.err != nil {
		return
	}

	if len(e.password) <= 100 {
		return
	}

	e.err = apperr.NewValidation("password", "length", "maximum is 100")
}
