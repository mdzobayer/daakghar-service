package conversation

import (
	dbdt "github.com/daakghar-service/data/db"
	"gopkg.in/mgo.v2/bson"
)

// Update handles conversation update api data
type Update struct {
	ClientID          string   `json:"ClientID"`
	ConversationName  string   `json:"ConversationName"`
	ConversationAdmin string   `json:"ConversationAdmin"`
	Members           []Member `json:"Members"`

	err error
}

// ToDB converts conversation create api data to DB data
func (u Update) ToDB(ID string, db *dbdt.Conversation) {
	db.Basic.ID = bson.ObjectIdHex(ID)
	db.ClientID = u.ClientID
	db.ConversationName = u.ConversationName
	db.ConversationAdmin = u.ConversationAdmin

	var dbMember dbdt.Member

	for _, member := range u.Members {
		memberToDB(member, &dbMember)
		db.Members = append(db.Members, dbMember)
	}
}

// Valid valids product data
func (u *Update) Valid() {

}

// Err returns product update package errors
func (u Update) Err() error {
	return u.err
}
