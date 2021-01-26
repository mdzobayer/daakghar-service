package user

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/daakghar-service/conn"

	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/data/api/filters"
	"github.com/daakghar-service/data/api/user"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/jwork"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

// admRead handles admin
type admRead struct {
	id      filters.ID
	apidt   user.AdmGet
	UsrInfo jwork.UsrInfo
	resp    io.Reader

	err error
}

// hasErr checks whether user package has any errors
func (a admRead) hasErr() bool {
	if a.err != nil {
		return true
	}

	return false
}

// fetchUser fetches user for user get
func (a *admRead) fetchUser() {
	if a.hasErr() {
		return
	}

	a.UsrInfo.FindUser()
	if a.UsrInfo.Err() != nil {
		a.err = errors.Wrap(a.UsrInfo.Err(), "jwork.user.admRead.fetchUser, could not find user")
	}
}

// fetchRole fetches role for admin user
func (a *admRead) fetchRole() {
	if a.hasErr() {
		return
	}

	a.UsrInfo.FindRole()
	if a.UsrInfo.Err() != nil {
		a.err = errors.Wrap(a.UsrInfo.Err(), "jwork.user.admRead.fetchUser, could not find role")
	}
}

// accessCheck checks whether admin can access detailed product get
func (a *admRead) accessCheck() {
	if a.hasErr() {
		return
	}

	if !a.UsrInfo.Role().UserAccess.CanRead() {
		a.err = errors.Wrap(apperr.NewAuthentication("check user access"), "jwork.product.admRead.accessCheck")
	}
}

// read reads admin user data from database
func (a *admRead) read() {
	if a.hasErr() {
		return
	}

	userQ := dbq.NewUser(conn.DB().DB())
	user, err := userQ.GetByID(bson.ObjectIdHex(a.id.Val))

	if err != nil {
		a.err = errors.Wrap(err, "jwork.user.admRead.read, db query")
		return
	}

	a.apidt.FromDB(user)
}

// genResp returns admin user info after a successful query in a user db
func (a *admRead) genResp() {
	if a.hasErr() {
		return
	}

	buf, err := json.Marshal(a.apidt)
	if err != nil {
		a.err = errors.Wrap(err, "jwork.user.admRead.genResp")
		return
	}

	a.resp = bytes.NewBuffer(buf)
}

// Work implements worker interface
func (a *admRead) Work() {
	a.fetchUser()
	a.fetchRole()
	a.accessCheck()
	a.read()
	a.genResp()
}

// Err returns admin user package errors
func (a *admRead) Err() error {
	return a.err
}

// Resp returns admin user response
func (a *admRead) Resp() io.Reader {
	return a.resp
}

// NewAdmGet returns worker to get admin user data
func NewAdmGet(user string, fltid filters.ID) jwork.Worker {
	return &admRead{
		id:      fltid,
		UsrInfo: jwork.UsrInfo{UserName: user},
	}
}
