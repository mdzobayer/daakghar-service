package user

import (
	"time"

	"github.com/daakghar-service/data/app"
	"github.com/daakghar-service/data/db"
	"gopkg.in/mgo.v2/bson"
)

// Update handles user update data
type Update struct {
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

// ToDB sets data to db
func (u Update) ToDB(ID string, db *db.User) {
	db.Basic.ID = bson.ObjectIdHex(ID)

	db.UserName = u.UserName.String()
	db.Name = u.Name
	db.Password, u.err = u.Password.Hash()
	//db.RoleID = bson.ObjectIdHex(u.RoleID)
	db.Email = u.Email
	db.Phone = u.Phone
	db.Address = u.Address

	db.DateOfBirth = u.DateOfBirth
	db.Gender = u.Gender
}

// Valid valids review data
func (u *Update) Valid() {

}

// Err returns review udpate package errors
func (u Update) Err() error {
	return u.err
}
