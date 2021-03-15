package filter

import (
	"github.com/daakghar-service/dbq"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userName struct {
	usrName string
}

func (u userName) Build(c *mgo.Collection) *mgo.Query {
	return c.Find(bson.M{
		"user_name": u.usrName,
	})
}

// NewUserName returns query builder find by user name
func NewUserName(usrName string) dbq.Builder {
	return &userName{usrName: usrName}
}
