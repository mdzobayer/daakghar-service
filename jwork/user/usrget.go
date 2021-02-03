package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/api/filters"
	"github.com/daakghar-service/data/api/user"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/jwork"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

// usrRead handles usr info read data
type usrRead struct {
	id      filters.ID
	apidt   user.UsrGet
	UsrInfo jwork.UsrInfo
	resp    io.Reader

	err error
}

// hasErr checks whether user package has any errors
func (u usrRead) hasErr() bool {
	if u.err != nil {
		return true
	}

	return false
}

// fetchUser fetches user for user get
func (u *usrRead) fetchUser() {
	if u.hasErr() {
		return
	}

	u.UsrInfo.FindUser()
	if u.UsrInfo.Err() != nil {
		u.err = errors.Wrap(u.UsrInfo.Err(), "jwork.user.usrRead.fetchUser, could not find user")
	}
}

// checkUser checks whether an user asked for his own user info
func (u *usrRead) checkUser() {
	if u.hasErr() {
		return
	}
	fmt.Println("\n\n apidt = ", u.apidt.UserName)
	fmt.Println("UsrInfo = ", u.UsrInfo.UserName)
	if u.apidt.UserName != u.UsrInfo.UserName {
		u.err = errors.Wrap(apperr.NewAuthentication("Unauthorized user data request"), "jwork.user.checkUser")
	}
}

// read reads user data from database
func (u *usrRead) read() {
	if u.hasErr() {
		return
	}
	userQ := dbq.NewUser(conn.DB().DB())
	user, err := userQ.GetByID(bson.ObjectIdHex(u.id.Val))
	if err != nil {
		u.err = errors.Wrap(err, "jwork.user.usrRead.read, db query")
		return
	}

	u.apidt.FromDB(user)
}

// genResp returns user info after a successful query in a user db
func (u *usrRead) genResp() {
	if u.hasErr() {
		return
	}

	buf, err := json.Marshal(u.apidt)
	if err != nil {
		u.err = errors.Wrap(err, "jwork.user.usrRead.genResp")
		return
	}

	u.resp = bytes.NewBuffer(buf)
}

// Work implements worker interface
func (u *usrRead) Work() {
	u.fetchUser()
	u.read()
	//u.checkUser()
	u.genResp()
}

// Err returns user package errors
func (u *usrRead) Err() error {
	return u.err
}

// Resp returns user response
func (u *usrRead) Resp() io.Reader {
	return u.resp
}

// NewUsrRead returns worker to get regular (non-admin ) user data
func NewUsrRead(user string, fltid filters.ID) jwork.Worker {
	return &usrRead{
		id:      fltid,
		UsrInfo: jwork.UsrInfo{UserName: user},
	}
}
