package user

import (
	"time"

	"github.com/pkg/errors"
	"github.com/daakghar-service/data/app"
	"github.com/daakghar-service/data/db"
	"gopkg.in/mgo.v2/bson"
)

// Create handles user create data
type Create struct {
	UserName    app.Username `json:"user_name"`
	Name        string       `json:"name"`
	Password    app.Password `json:"password"`
	RoleID      string       `json:"role_id"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	Address     []string     `json:"address"`
	DateOfBirth time.Time    `json:"date_of_birth"`
	Gender      string       `json:"gender"`

	err error
}

// hasErr checks whether user api has any errors
func (c Create) hasErr() bool {
	if c.err != nil {
		return true
	}

	return false
}

// usrNameValid checks whether username is valid
func (c *Create) usrNameValid() {
	if c.hasErr() {
		return
	}

	err := c.UserName.Valid()
	if err != nil {
		c.err = errors.Wrap(err, "data.api.user.create.usrNameValid")
	}
}

// passValid checks whether password is valid
func (c *Create) passValid() {
	if c.hasErr() {
		return
	}

	err := c.Password.Valid()
	if err != nil {
		c.err = errors.Wrap(err, "data.api.user.create.passValid")
	}
}

// Valid valids user data
func (c *Create) Valid() {
	c.usrNameValid()
	c.passValid()
}

// Err returns error values
func (c Create) Err() error {
	return c.err
}

// ToDB sets data to db.User
func (c Create) ToDB(db *db.User) {
	db.UserName = c.UserName.String()
	db.Name = c.Name
	db.Password, c.err = c.Password.Hash()
	if c.RoleID != "" {
		db.RoleID = bson.ObjectIdHex(c.RoleID)
	}
	db.Email = c.Email
	db.Phone = c.Phone
	db.Address = c.Address

	db.Verified = false

	db.DateOfBirth = c.DateOfBirth
	db.Gender = c.Gender
}
