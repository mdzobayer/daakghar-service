package dbq

import (
	"gopkg.in/mgo.v2"
)

// Builder wraps query build operations
type Builder interface {
	Build(c *mgo.Collection) *mgo.Query
}
