package dbq

import (
	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/data/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Conversation handles conversation db queries
type Conversation struct {
	db *mgo.Database
}

// collection returns the collection that is holding the conversation data
func (c *Conversation) collection() *mgo.Collection {
	return c.db.C(conversationC)
}

// GetByID returns conversation filtered by _id
func (c *Conversation) GetByID(id bson.ObjectId) (dt db.Conversation, err error) {

	err = c.collection().FindId(id).One(&dt)

	if err == mgo.ErrNotFound {
		err = apperr.NewNotFound("product", id.Hex())
	}

	return
}

// Put creates or updates a conversation
func (c *Conversation) Put(dt *db.Conversation) (err error) {
	dt.PreparePut()

	_, err = c.collection().UpsertId(dt.ID, bson.M{"$set": dt})

	return
}

// Delete deletes conversation
func (c *Conversation) Delete(dt *db.Conversation) (err error) {
	dt.PrepareDelete()

	err = c.Put(dt)

	return
}

// NewConversation returns conversation
func NewConversation(db *mgo.Database) *Conversation {
	return &Conversation{
		db: db,
	}
}
