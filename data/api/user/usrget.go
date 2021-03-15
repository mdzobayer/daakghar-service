package user

import (
	"time"

	"github.com/daakghar-service/data/db"
)

// UsrGet handles user get data
type UsrGet struct {
	UserName    string    `json:"user_name"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     []string  `json:"address,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
}

// FromDB sets data from db
func (g *UsrGet) FromDB(db db.User) {

	g.UserName = db.UserName
	g.Name = db.Name
	g.Email = db.Email
	g.Phone = db.Phone
	g.Address = db.Address
	g.DateOfBirth = db.DateOfBirth
	g.Gender = db.Gender
}
