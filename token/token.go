package token

import (
	"strings"
	"time"

	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/token/aes"
	"github.com/pkg/errors"
)

// User struct
type User struct {
	UserName string
	Expire   time.Time

	err error
}

// errF returns the user struct error
func (u User) errF() error {
	return u.err
}

// valid checks the validity of a token
func (u User) valid() {

	if u.errF() != nil {
		return
	}

	var s = u.string()

	var tmp = strings.Split(s, "@")

	if len(tmp) == 2 {
		return
	}
	u.err = apperr.NewValidation("field", "length", "equal to two")
}

// string converts a user struct to a string
func (u User) string() string {
	if u.errF() != nil {
		return ""
	}
	return u.UserName + "@" + u.Expire.Format(time.RFC1123Z)
}

// fromString converts a token to a user struct
func (u *User) fromString(token string) {

	if u.errF() != nil {
		return
	}

	s := strings.Split(token, "@")

	u.UserName = s[0]
	u.Expire, u.err = time.Parse(time.RFC1123Z, s[1])

	return
}

// Encrypt encrpyts a user to a token
func Encrypt(u User, key string) (token string, err error) {

	if token, err = aes.Encrypt(u.string(), key); err != nil {
		u.err = errors.Wrap(err, "token.Encrypt, Unable to encrypt a user data")
		return
	}

	return
}

// Decrypt decrypts a user from a request token
func Decrypt(token string, key string) (u User) {

	var decodeToken string
	var err error
	decodeToken, err = aes.Decrypt(token, key)

	if err != nil {
		u.err = errors.Wrap(err, "token.Decrypt, Unable to decrypt a user data")
		return
	}

	u.fromString(decodeToken)

	return
}
