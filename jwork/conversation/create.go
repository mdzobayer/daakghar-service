package conversation

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/api"
	"github.com/daakghar-service/data/api/conversation"
	"github.com/daakghar-service/data/db"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/jwork"
	"github.com/pkg/errors"
)

// create stores conversation api data
type create struct {
	Dt      conversation.Create
	UsrInfo jwork.UsrInfo
	err     error
	id      api.ID
	resp    io.Reader
}

// fetchUser fetches user (an admin) who has made the conversation create request
func (c *create) fetchUser() {
	if c.hasErr() {
		return
	}

	c.UsrInfo.FindUser()
	if c.UsrInfo.Err() != nil {
		c.err = errors.Wrap(c.UsrInfo.Err(), "jwork.conversation.create.fetchUser, could not find user")
	}
}

func (c *create) fetchRole() {
	if c.hasErr() {
		return
	}

	c.UsrInfo.FindRole()
	if c.UsrInfo.Err() != nil {
		c.err = errors.Wrap(c.UsrInfo.Err(), "jwork.conversation.create.fetchUser, could not find role")
		return
	}
}

// accessCheck checks wheather the current user has the right to create conversation
func (c *create) accessCheck() {
	if c.hasErr() {
		return
	}

	if !c.UsrInfo.Role().ConversationAccess.CanWrite() {
		c.err = errors.Wrap(apperr.NewAuthentication("write conversation"), "jwork.conversation.create.accessCheck")
	}
}

// valid validates conversation payload
func (c *create) valid() {
	if c.hasErr() {
		return
	}

	c.Dt.Valid()
	if c.Dt.Err() != nil {
		c.err = errors.Wrap(c.Dt.Err(), "jwork.user.create.valid")
	}
}

// hasErr checks whether conversation createion has error
func (c create) hasErr() bool {
	if c.err != nil {
		return true
	}

	return false
}

// put creates or update converstion
func (c *create) put() {
	if c.hasErr() {
		return
	}

	conversation := dbq.NewConversation(conn.DB().DB())

	dbdt := db.Conversation{}
	c.Dt.ToDB(&dbdt)

	err := conversation.Put(&dbdt)
	if err != nil {
		c.err = errors.Wrap(err, "jwork.conversation.create.put. could not create conversation")
		return
	}

	c.id.FromDB(dbdt.Basic)
}

// genResp creates api response for conversation
func (c *create) genResp() {
	if c.hasErr() {
		return
	}

	b, err := json.Marshal(c.id)
	if err != nil {
		c.err = errors.Wrap(err, "jwork.conversation.create,genResp, could not marshal id object")
		return
	}

	c.resp = bytes.NewBuffer(b)
}

// Work implements Worker interface for jwork
func (c *create) Work() {
	c.valid()
	c.fetchUser()
	c.fetchRole()
	c.accessCheck()
	c.put()
	c.genResp()
}

// Err returns jwork conversation errors
func (c create) Err() error {
	return c.err
}

// Resp returns jwork conversation response
func (c create) Resp() io.Reader {
	return c.resp
}

// NewCreate creates jworker for conversation
func NewCreate(userName string, dt conversation.Create) jwork.Worker {
	return &create{
		Dt:      dt,
		UsrInfo: jwork.UsrInfo{UserName: userName},
	}
}
