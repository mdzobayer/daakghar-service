package db

import (
	"time"

	"github.com/daakghar-service/data/app"
	"gopkg.in/mgo.v2/bson"
)

// User holds User data
type User struct {
	Basic       `bson:",inline"`
	UserName    string        `bson:"user_name,omitempty"`
	Name        string        `bson:"name,omitempty"`
	Password    string        `bson:"password,omitempty"`
	RoleID      bson.ObjectId `bson:"role_id,omitempty"`
	Email       string        `bson:"email,omitempty"`
	Verified    bool          `bson:"verified,omitempty"`
	Phone       string        `bson:"phone,omitempty"`
	Address     []string      `bson:"address,omitempty"`
	DateOfBirth time.Time     `bson:"data_of_birth,omitempty"`
	Gender      string        `bson:"gender,omitempty"`
}

// Valid returns errors for invalid admin data
func (u User) Valid() error {
	return nil
}

// AdminUser returns admin user
func AdminUser(roleID bson.ObjectId) (User, error) {
	pass := app.Password("adminadmin")
	hpass, err := pass.Hash()
	if err != nil {
		return User{}, err
	}

	return User{
		UserName: "admin",
		Name:     "Admin",
		Password: hpass,
		RoleID:   roleID,
		Verified: true,
	}, nil
}
