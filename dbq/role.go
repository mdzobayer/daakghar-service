package dbq

import (
	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/data/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Role struct
type Role struct {
	db *mgo.Database
}

func (r *Role) collection() *mgo.Collection {
	return r.db.C(roleC)
}

// GetByID retruns
func (r *Role) GetByID(id bson.ObjectId) (dt db.Role, err error) {

	err = r.collection().FindId(id).One(&dt)

	if err == mgo.ErrNotFound {
		err = apperr.NewNotFound("role", id.Hex())
	}

	return
}

// Put prepar
func (r *Role) Put(dt *db.Role) (err error) {

	dt.PreparePut()

	_, err = r.collection().UpsertId(dt.ID, bson.M{"$set": dt})

	return
}

// Delete returns
func (r *Role) Delete(dt db.Role) (err error) {
	dt.PrepareDelete()

	err = r.Put(&dt)

	return
}

// NewRole returns new role
func NewRole(db *mgo.Database) *Role {
	return &Role{
		db: db,
	}
}
