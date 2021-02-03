package conn

import (
	"github.com/daakghar-service/config"
	"gopkg.in/mgo.v2"
)

// DBConnector wraps database conncetion
type DBConnector interface {
	Connect(conf config.Mgo)
	Err() error
	DB() *mgo.Database
}
