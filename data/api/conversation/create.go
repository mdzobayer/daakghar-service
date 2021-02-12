package conversation

import (
	"time"

	dbdt "github.com/daakghar-service/data/db"
)

// Create holds conversation api data
type Create struct {
	ClientID          string   `json:"ClientID"`
	ConversationName  string   `json:"ConversationName"`
	ConversationAdmin string   `json:"ConversationAdmin"`
	Members           []Member `json:"Members"`

	err error
}

// Member holds conversation member info
type Member struct {
	MemberID string `json:"MemberID"`
	NickName string `json:"NickName"`
	Role     string `json:"Role"`

	JoinedBy    string    `json:"JoinedBy"`
	JoiningTime time.Time `json:"JoiningTime"`

	RemovedBy    string    `json:"RemovedBy"`
	RemovingTime time.Time `json:"RemovingTime"`

	LeavingTime time.Time `json:"LeavingTime"`
}

// Valid valids conversation data
func (c *Create) Valid() {

}

// Err returns conversation create package errors
func (c Create) Err() error {
	return c.err
}

// ToDB converts conversation create api data to DB data
func (c Create) ToDB(db *dbdt.Conversation) {
	db.ClientID = c.ClientID
	db.ConversationName = c.ConversationName
	db.ConversationAdmin = c.ConversationAdmin

	var dbMember dbdt.Member

	for _, member := range c.Members {
		memberToDB(member, &dbMember)
		db.Members = append(db.Members, dbMember)
	}
}

// MemberToDB converts conversation create members api data to DB data
func memberToDB(jMember Member, bMember *dbdt.Member) {
	bMember.JoinedBy = jMember.JoinedBy
	bMember.JoiningTime = jMember.JoiningTime
	bMember.LeavingTime = jMember.LeavingTime
	bMember.MemberID = jMember.MemberID
	bMember.NickName = jMember.NickName
	bMember.RemovedBy = jMember.RemovedBy
	bMember.RemovingTime = jMember.RemovingTime
	bMember.Role = jMember.Role
}
