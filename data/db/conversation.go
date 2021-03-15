package db

import (
	"time"
)

// Conversation holds conversation basic info
type Conversation struct {
	Basic             `bson:",inline"`
	ClientID          string   `bson:"ClientID"`
	ConversationName  string   `bson:"ConversationName"`
	ConversationAdmin string   `bson:"ConversationAdmin"`
	Members           []Member `bson:"Members"`
}

// Member holds conversation member info
type Member struct {
	MemberID string `bson:"MemberID"`
	NickName string `bson:"NickName"`
	Role     string `bson:"Role"`

	JoinedBy    string    `bson:"JoinedBy"`
	JoiningTime time.Time `bson:"JoiningTime"`

	RemovedBy    string    `bson:"RemovedBy,omitempty"`
	RemovingTime time.Time `bson:"RemovingTime,omitempty"`

	LeavingTime time.Time `bson:"LeavingTime,omitempty"`
}
