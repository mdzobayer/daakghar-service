package user

import (
	"github.com/daakghar-service/data/api"
	"github.com/daakghar-service/data/app"
	"github.com/daakghar-service/data/db"
)

// AdmGet handles user get data
type AdmGet struct {
	api.Basic
	UserName app.Username `json:"user_name"`
	Name     string       `json:"name"`
	RoleID   string       `json:"role_id,omitempty"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	Address  []string     `json:"address,omitempty"`
	Verified bool         `json:"verified"`
}

// FromDB sets data from db
func (g *AdmGet) FromDB(db db.User) {
	g.Basic.FromDB(db.Basic)

	g.Name = db.Name
	if db.RoleID != "" {
		g.RoleID = db.RoleID.Hex()
	}
	g.Email = db.Email
	g.Phone = db.Phone
	g.Address = db.Address
	g.Verified = db.Verified
}
