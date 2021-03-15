package api

import (
	"time"

	"github.com/daakghar-service/data/db"
	"gopkg.in/mgo.v2/bson"
)

// ID handles api id data
type ID struct {
	ID string `json:"id"`
}

// ToDB sets api data to db data
func (i ID) ToDB(db *db.Basic) {
	db.ID = bson.ObjectIdHex(i.ID)
}

// FromDB sets db data to api data
func (i *ID) FromDB(db db.Basic) {
	i.ID = db.ID.Hex()
}

// Basic handles api basic data
type Basic struct {
	ID

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
	DeletedAt time.Time `json:"deleted_at"`
}

// ToDB sets api data to db data
func (b Basic) ToDB(db *db.Basic) {
	b.ID.ToDB(db)

	db.CreatedAt = b.CreatedAt
	db.UpdatedAt = b.UpdatedAt
	db.Deleted = b.Deleted
}

// FromDB sets db data to api data
func (b *Basic) FromDB(db db.Basic) {
	b.ID.FromDB(db)

	b.CreatedAt = db.CreatedAt
	b.UpdatedAt = db.UpdatedAt
	b.Deleted = db.Deleted
}
