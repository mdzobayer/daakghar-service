package conversation

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/api/conversation"
	"github.com/daakghar-service/data/api/filters"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/jwork"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type privateRead struct {
	id filters.ID

	apidt conversation.Get

	resp io.Reader

	err error
}

func (c privateRead) hasErr() bool {
	if c.err != nil {
		return true
	}

	return false
}

func (c *privateRead) read() {
	if c.hasErr() {
		return
	}

	conversation := dbq.NewConversation(conn.DB().DB())
	conversn, err := conversation.GetByID(bson.ObjectIdHex(c.id.Val))

	if err != nil {
		c.err = errors.Wrap(err, "jwork.conversatoin.privateRead.read, db query")
		return
	}

	c.apidt.FromDB(conversn)
}

func (c *privateRead) genResp() {
	if c.hasErr() {
		return
	}

	b, err := json.Marshal(c.apidt)

	if err != nil {
		c.err = errors.Wrap(err, "jwork.conversation.privateRead.genResp, could not marshal")
		return
	}

	c.resp = bytes.NewBuffer(b)
}

func (c *privateRead) Work() {
	c.read()
	c.genResp()
}

func (c *privateRead) Err() error {
	return c.err
}

func (c *privateRead) Resp() io.Reader {
	return c.resp
}

// NewPrivateRead returns private conversation data
func NewPrivateRead(flt filters.ID) jwork.Worker {
	return &privateRead{
		id: flt,
	}
}
