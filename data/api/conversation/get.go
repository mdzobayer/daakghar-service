package conversation

import (
	"github.com/daakghar-service/data/api"
	"github.com/daakghar-service/data/db"
)

// Get handles conversation get data
type Get struct {
	api.Basic

	ClientID          string `json:"ClientID"`
	ConversationName  string `json:"ConversationName"`
	ConversationAdmin string `json:"ConversationAdmin"`
}

// FromDB sets conversation data from DB
func (c *Get) FromDB(db db.Conversation) {
	c.Basic.FromDB(db.Basic)

	c.ClientID = db.ClientID
	c.ConversationAdmin = db.ConversationAdmin
	c.ConversationName = db.ConversationName
}
