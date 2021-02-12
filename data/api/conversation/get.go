package conversation

import (
	"github.com/daakghar-service/data/api"
	dbdt "github.com/daakghar-service/data/db"
)

// Get handles conversation get data
type Get struct {
	api.Basic

	ClientID          string   `json:"ClientID"`
	ConversationName  string   `json:"ConversationName"`
	ConversationAdmin string   `json:"ConversationAdmin"`
	Members           []Member `json:"Members"`
}

// FromDB sets conversation data from DB
func (c *Get) FromDB(db dbdt.Conversation) {
	c.Basic.FromDB(db.Basic)

	c.ClientID = db.ClientID
	c.ConversationAdmin = db.ConversationAdmin
	c.ConversationName = db.ConversationName

	var jMember Member

	for _, member := range db.Members {
		memberFromDB(member, &jMember)
		c.Members = append(c.Members, jMember)
	}

}

// MemberFromDB converts conversation create members api data to DB data
func memberFromDB(bMember dbdt.Member, jMember *Member) {
	jMember.JoinedBy = bMember.JoinedBy
	jMember.JoiningTime = bMember.JoiningTime
	jMember.LeavingTime = bMember.LeavingTime
	jMember.MemberID = bMember.MemberID
	jMember.NickName = bMember.NickName
	jMember.RemovedBy = bMember.RemovedBy
	jMember.RemovingTime = bMember.RemovingTime
	jMember.Role = bMember.Role
}
