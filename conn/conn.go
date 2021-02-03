package conn

import (
	"github.com/daakghar-service/config"
	"gopkg.in/mgo.v2"
)

type database struct {
	db  *mgo.Database
	err error
}

// Connect connects to mongodb
func (d *database) Connect(conf config.Mgo) {
	s, err := mgo.Dial(conf.URI)

	if err != nil {
		d.err = err
		return
	}

	d.db = s.DB("")
}

// Err returns error value
func (d database) Err() error {
	return d.err
}

// DB returns mongodb database connection
func (d database) DB() *mgo.Database {
	return d.db
}

var db database

// DB returns DBConnector
func DB() DBConnector {
	return &db
}
