package app

import (
	"fmt"
	"regexp"

	"github.com/daakghar-service/apperr"
)

// Username represents user name
type Username string

// String converts Username to string
func (u Username) String() string {
	return string(u)
}

// Valid valids user password
func (u Username) Valid() error {

	eu := validUser{
		username: u,
	}

	eu.minLenValid()
	eu.maxLenValid()
	eu.charsValid()

	return eu.Err()
}

// validUser handles username validity
type validUser struct {
	username Username
	err      error
}

// String returns validUser string
func (e *validUser) String() string {
	return fmt.Sprintf("Username=%s, error=%s", e.username, e.err)
}

// Err returns validUser error
func (e *validUser) Err() error {
	return e.err
}

// minLenValid checks for minimum length validity of a username
func (e *validUser) minLenValid() {
	if e.err != nil {
		return
	}

	if len(e.username) >= 1 {
		return
	}

	e.err = apperr.NewValidation("username", "length", "minimum is 1")
}

// maxLenValid checks for maximum length validity of a username
func (e *validUser) maxLenValid() {
	if e.err != nil {
		return
	}

	if len(e.username) <= 100 {
		return
	}

	e.err = apperr.NewValidation("username", "length", "maximum is 100")
}

// charsValid checks for valid characters in a username
func (e *validUser) charsValid() {

	if e.err != nil {
		return
	}

	e.matchPattern()
}

// matchPattern checks for valid username regular expression patterns
func (e *validUser) matchPattern() {
	var matched bool

	matched = e.validateName()

	if matched {
		return
	}

	e.err = apperr.NewValidation("username", "invalid characters", "May contain only letter, digits or underscores")
}

// validateName compiles regex rules and test a username against it
func (e *validUser) validateName() (matched bool) {

	var validName = regexp.MustCompile(`^\w+$`)

	matched = validName.MatchString(string(e.username))

	return
}

// detailed validation rule needs to be written later
