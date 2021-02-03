package conversation

import "github.com/daakghar-service/data/db"

// Create holds conversation api data
type Create struct {
	ClientID          string `json:"ClientID"`
	ConversationName  string `json:"ConversationName"`
	ConversationAdmin string `json:"ConversationAdmin"`

	err error
}

// Valid valids conversation data
func (c *Create) Valid() {

}

// Err returns conversation create package errors
func (c Create) Err() error {
	return c.err
}

// ToDB converts conversation create api data to DB data
func (c Create) ToDB(db *db.Conversation) {
	db.ClientID = c.ClientID
	db.ConversationName = c.ConversationName
	db.ConversationAdmin = c.ConversationAdmin
}
