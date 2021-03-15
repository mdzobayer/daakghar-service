package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/config"
	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/api/user"
	"github.com/daakghar-service/data/db"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/dbq/filter"
	"github.com/daakghar-service/hash"
	"github.com/daakghar-service/jwork"
	tk "github.com/daakghar-service/token"
	"github.com/pkg/errors"
)

// token stores api data for token creation
type token struct {
	apidt user.TokenReq

	dbdt  db.User
	usrtk user.TokenResp

	resp io.Reader

	err error
}

// hasErr checks whether token worker has errors
func (t *token) hasErr() bool {
	if t.err != nil {
		return true
	}

	return false
}

// dbFind searches for username
func (t *token) dbFind() {
	if t.hasErr() {
		return
	}

	flt := filter.NewUserName(t.apidt.UserName.String())
	usrs, err := dbq.NewUser(conn.DB().DB()).List(flt)
	if err != nil {
		t.err = errors.Wrap(err, "jwork.user.token.dbFind, db query")
		return
	}

	if len(usrs) < 1 {
		t.err = errors.Wrap(
			apperr.NewNotFound("user", t.apidt.UserName.String()),
			"jwork.user.token.dbFind, user not found",
		)
		return
	}

	t.dbdt = usrs[0]
}

// valid checks whether token api data is valid
func (t *token) valid() {
	if t.hasErr() {
		return
	}

	t.apidt.Valid()
	if t.apidt.Err() != nil {
		t.err = errors.Wrap(t.apidt.Err(), "jwork.user.token.valid")
	}
}

// passCheck checks whether the user password is valid
func (t *token) passCheck() {
	if t.hasErr() {
		return
	}

	if err := hash.CompareHashAndPassword(t.dbdt.Password, string(t.apidt.Password)); err != nil {
		t.err = errors.Wrap(apperr.NewNotFound("user", "wrong password"), "jwork.user.token.passCheck")
	}
}

// genToken generates token from username and password
func (t *token) genToken() {
	if t.hasErr() {
		return
	}

	usr := tk.User{
		UserName: t.dbdt.UserName,
		Expire:   time.Now().UTC().Add(time.Hour * 24 * 30), // Expiry date for a token is set to a month
	}

	// fmt.Println("token user", usr, config.Get().Token().Key)
	usrtk, err := tk.Encrypt(usr, config.Get().Token().Key)
	if err != nil {
		t.err = errors.Wrap(err, "jwork.user.token.genToken")
		return
	}
	fmt.Println("token", usrtk)

	t.usrtk.Token = usrtk
}

// genResp generates response for token worker
func (t *token) genResp() {
	if t.hasErr() {
		return
	}

	b, err := json.Marshal(t.usrtk)
	if err != nil {
		t.err = errors.Wrap(err, "jwork.user.token.genResp, could not marshal roken resp object")
		return
	}

	t.resp = bytes.NewBuffer(b)
}

// Work implements worker interface
func (t *token) Work() {
	t.valid()
	t.dbFind()
	t.passCheck()
	t.genToken()
	t.genResp()
}

// Err returns token worker error
func (t token) Err() error {
	return t.err
}

// Resp returns token worker response
func (t token) Resp() io.Reader {
	return t.resp
}

// NewToken retuns token worker
func NewToken(dt user.TokenReq) jwork.Worker {
	return &token{
		apidt: dt,
	}
}
