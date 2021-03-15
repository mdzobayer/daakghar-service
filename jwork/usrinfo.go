package jwork

import (
	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/db"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/dbq/filter"
	"github.com/pkg/errors"
)

// UsrInfo provides user and role info
type UsrInfo struct {
	UserName string

	user *db.User
	role *db.Role

	err error
}

// User returns db user
func (u UsrInfo) User() *db.User {
	return u.user
}

// Role returns db role
func (u UsrInfo) Role() *db.Role {
	return u.role
}

// FindUser finds user from db
func (u *UsrInfo) FindUser() {
	if u.hasErr() {
		return
	}

	du := dbq.NewUser(conn.DB().DB())

	usrs, err := du.List(filter.NewUserName(u.UserName))

	if err != nil {
		u.err = errors.Wrap(err, "jwork.UsrInfo.FindUser")
		return
	}

	if len(usrs) == 0 {
		u.err = apperr.NewNotFound("user", u.UserName)
		return
	}

	u.user = &usrs[0]
}

// FindRole finds role from db
func (u *UsrInfo) FindRole() {
	if u.hasErr() {
		return
	}

	if u.user.RoleID == "" {
		u.err = apperr.NewNotFound("user role", "empty role")
		return
	}

	dr := dbq.NewRole(conn.DB().DB())
	rl, err := dr.GetByID(u.user.RoleID)

	if err != nil {
		u.err = errors.Wrap(err, "jwork,UsrInfo.FindRole")
		return
	}

	u.role = &rl
}

// Err returns error value
func (u UsrInfo) Err() error {
	return u.err
}

// hasErr checks error
func (u UsrInfo) hasErr() bool {
	if u.err != nil {
		return true
	}

	return false
}
