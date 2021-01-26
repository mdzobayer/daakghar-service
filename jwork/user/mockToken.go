package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/daakghar-service/config"
	"github.com/daakghar-service/data/api/user"
	"github.com/daakghar-service/data/db"
	"github.com/daakghar-service/jwork"
	tk "github.com/daakghar-service/token"
	"github.com/pkg/errors"
)

// mockToken stores api data for mockToken creation
type mockToken struct {
	apidt user.TokenReq

	dbdt  db.User
	usrtk user.TokenResp

	resp io.Reader

	err error
}

// hasErr checks whether mockToken worker has errors
func (t *mockToken) hasErr() bool {
	if t.err != nil {
		return true
	}

	return false
}

// dbFind searchs for username
func (t *mockToken) dbFind() {
	if t.hasErr() {
		return
	}

	t.dbdt.UserName = "TestName"
}

// valid checks whether mockToken api data is valid
func (t *mockToken) valid() {
	if t.hasErr() {
		return
	}

	t.apidt.Valid()
	if t.apidt.Err() != nil {
		t.err = errors.Wrap(t.apidt.Err(), "jwork.user.mockToken.valid")
	}
}

// passCheck checks whether the user password is valid
func (t *mockToken) passCheck() {
	if t.hasErr() {
		return
	}

}

// genmockToken generates mockToken from username and password
func (t *mockToken) genToken() {
	if t.hasErr() {
		return
	}

	usr := tk.User{
		UserName: t.dbdt.UserName,
		Expire:   time.Now().UTC().Add(time.Hour * 24 * 30), // Expiry date for a mockToken is set to a month
	}

	//fmt.Println("mockToken user", usr, config.Get().mockToken().Key)
	usrtk, err := tk.Encrypt(usr, config.Get().Token().Key)
	if err != nil {
		t.err = errors.Wrap(err, "jwork.user.mockToken.genToken")
		return
	}
	fmt.Println("mockToken", usrtk)

	t.usrtk.Token = usrtk
}

// genResp generates response for mockToken worker
func (t *mockToken) genResp() {
	if t.hasErr() {
		return
	}

	b, err := json.Marshal(t.usrtk)
	if err != nil {
		t.err = errors.Wrap(err, "jwork.user.mockToken.genResp, could not marshal roken resp object")
		return
	}

	t.resp = bytes.NewBuffer(b)
}

// Work implements worker interface
func (t *mockToken) Work() {
	t.valid()
	t.dbFind()
	t.passCheck()
	t.genToken()
	t.genResp()
}

// Err returns mockToken worker error
func (t mockToken) Err() error {
	return t.err
}

// Resp returns mockToken worker response
func (t mockToken) Resp() io.Reader {
	return t.resp
}

// NewMockToken retuns token worker
func NewMockToken(dt user.TokenReq) jwork.Worker {
	return &mockToken{
		apidt: dt,
	}
}
