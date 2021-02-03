package dbq

import (
	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/data/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User handles user db queries
type User struct {
	db *mgo.Database
}

func (u *User) collection() *mgo.Collection {
	return u.db.C(userC)
}

// Put creates or updates user data
func (u *User) Put(dt *db.User) (err error) {
	dt.PreparePut()

	_, err = u.collection().UpsertId(dt.ID, bson.M{"$set": dt})

	return
}

// GetByID returns user filtered by id
func (u *User) GetByID(id bson.ObjectId) (dt db.User, err error) {

	err = u.collection().FindId(id).One(&dt)

	if err == mgo.ErrNotFound {
		err = apperr.NewNotFound("admin", id.Hex())
	}

	return
}

// Delete deletes user
func (u *User) Delete(dt *db.User) (err error) {
	dt.PrepareDelete()

	err = u.Put(dt)

	return
}

// List fetches user list using Builder
func (u *User) List(b Builder) ([]db.User, error) {
	usrs := []db.User{}

	err := b.Build(u.collection()).All(&usrs)

	return usrs, err
}

// NewUser returns User
func NewUser(db *mgo.Database) *User {
	return &User{
		db: db,
	}
}
