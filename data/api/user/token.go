package user

import (
	"github.com/daakghar-service/data/app"
	"github.com/pkg/errors"
)

// TokenReq handles user token request
type TokenReq struct {
	UserName app.Username `json:"user_name"`
	Password app.Password `json:"password"`

	err error
}

// TokenResp handles user token response
type TokenResp struct {
	Token string `json:"token"`
}

// hasErr returns whether TokenReq has any error
func (t TokenReq) hasErr() bool {
	if t.err != nil {
		return true
	}

	return false
}

// nameValid checks whether a given username is valid
func (t *TokenReq) nameValid() {
	if t.hasErr() {
		return
	}

	if err := t.UserName.Valid(); err != nil {
		t.err = errors.Wrap(err, "data.user.Token.nameValid")
	}
}

// passValid checks whether a password is valid
func (t *TokenReq) passValid() {
	if t.hasErr() {
		return
	}

	if err := t.Password.Valid(); err != nil {
		t.err = errors.Wrap(err, "data.user.Token.passValid")
	}
}

// Valid valids Token request
func (t *TokenReq) Valid() {
	t.nameValid()
	t.passValid()
}

// Err returns error value
func (t TokenReq) Err() error {
	return t.err
}
