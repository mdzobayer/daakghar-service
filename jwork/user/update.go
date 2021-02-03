package user

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/api"
	"github.com/daakghar-service/data/api/filters"
	"github.com/daakghar-service/data/api/user"
	dbUser "github.com/daakghar-service/data/db"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/jwork"
	"github.com/pkg/errors"
)

// update handles user update request
type update struct {
	id filters.ID

	apiDt user.Update
	apiID api.ID

	resp io.Reader

	err error
}

// hasErr checks weather public product data package has errors
func (u update) hasErr() bool {
	if u.err != nil {
		return true
	}

	return false
}

// valid checks wheatehr user update api data valid
func (u *update) valid() {
	if u.hasErr() {
		return
	}

	u.apiDt.Valid()
	if u.apiDt.Err() != nil {
		u.err = errors.Wrap(u.apiDt.Err(), "jwork.user.update.valid")
	}
}

// put create or update a review
func (u *update) put() {
	if u.hasErr() {
		return
	}

	usr := dbq.NewUser(conn.DB().DB())

	dbdt := dbUser.User{}

	u.apiDt.ToDB(u.id.Val, &dbdt)

	if u.apiDt.Err() != nil {
		u.err = errors.Wrap(u.apiDt.Err(), "jwork.user.update.put, could not copy to db")
		return
	}

	err := usr.Put(&dbdt)
	if err != nil {
		u.err = errors.Wrap(err, "jwork.user.update.put, could not update user")
	}

	u.apiID.FromDB(dbdt.Basic)
}

// genResp creates response for public user data
func (u *update) genResp() {
	if u.hasErr() {
		return
	}

	b, err := json.Marshal(u.apiDt)
	if err != nil {
		u.err = errors.Wrap(err, "jwork.user.Update.genResp, could not marshal")
		return
	}

	u.resp = bytes.NewBuffer(b)
}

// Work implemnts worker interface for public user
func (u *update) Work() {
	u.valid()
	u.put()
	u.genResp()
}

// Err returns public user data package errors
func (u *update) Err() error {
	return u.err
}

// Resp returns response of the user throught io.Reader
func (u *update) Resp() io.Reader {
	return u.resp
}

// NewUpdate returns worker to update user for public
func NewUpdate(flt filters.ID, dt user.Update) jwork.Worker {
	return &update{
		id:    flt,
		apiDt: dt,
	}
}
