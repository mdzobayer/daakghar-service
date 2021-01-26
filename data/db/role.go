package db

import "github.com/daakghar-service/data/app"

// Role holds role data
type Role struct {
	Basic              `bson:",inline"`
	Name               string         `bson:"name"`
	RoleAccess         app.AccessList `bson:"role_access"`
	UserAccess         app.AccessList `bson:"user_access"`
	CredentialAccess   app.AccessList `bson:"credential_access"`
	ConversationAccess app.AccessList `bson:"conversation_access"`
}

// Valid returns errors for invalid role data
func (r Role) Valid() error {
	return nil
}

// AdminRole returns role with super access
func AdminRole() Role {
	return Role{
		Name:               "admin",
		RoleAccess:         app.AllAccesses(),
		UserAccess:         app.AllAccesses(),
		CredentialAccess:   app.AllAccesses(),
		ConversationAccess: app.AllAccesses(),
	}
}
