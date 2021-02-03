package db

// Conversation holds conversation basic info
type Conversation struct {
	Basic             `bson:",inline"`
	ClientID          string `bson:"ClientID"`
	ConversationName  string `bson:"ConversationName"`
	ConversationAdmin string `bson:"ConversationAdmin"`
}
