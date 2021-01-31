package user

import (
	"bytes"
	"encoding/json"
	"io"
	"net/smtp"

	"github.com/daakghar-service/conn"
	"github.com/daakghar-service/data/api"
	"github.com/daakghar-service/data/api/user"
	"github.com/daakghar-service/data/db"
	"github.com/daakghar-service/dbq"
	"github.com/daakghar-service/jwork"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// create stores user create data
type create struct {
	apiDt user.Create
	err   error

	id   api.ID
	resp io.Reader
}

// valid checks whether user.Create api data is valid
func (c *create) valid() {
	if c.hasErr() {
		return
	}

	c.apiDt.Valid()
	if c.apiDt.Err() != nil {
		c.err = errors.Wrap(c.apiDt.Err(), "jwork.user.create.valid")
	}
}

// hasErr checks whether jwork user package has any errors
func (c *create) hasErr() bool {
	if c.err != nil {
		return true
	}

	return false
}

// put creates or updates a user
func (c *create) put() {
	if c.hasErr() {
		return
	}

	usr := dbq.NewUser(conn.DB().DB())

	dbdt := db.User{}

	c.apiDt.ToDB(&dbdt)
	if c.apiDt.Err() != nil {
		c.err = errors.Wrap(c.apiDt.Err(), "jwork.user.create.put, could not copy to db")
		return
	}

	err := usr.Put(&dbdt)
	if err != nil {
		c.err = errors.Wrap(err, "jwork.user.create.put, could not create user")
		return
	}

	c.id.FromDB(dbdt.Basic)
}

// getResp generates response
func (c *create) getResp() {
	if c.hasErr() {
		return
	}

	b, err := json.Marshal(c.id)
	if err != nil {
		c.err = errors.Wrap(err, "jwork.customer.create.getResp, could not marshal id object")
		return
	}

	c.resp = bytes.NewBuffer(b)
}

// Work implements worker interface
func (c *create) Work() {
	c.valid()
	c.put()
	c.SendVerificationMail()
	c.getResp()
}

// Err returns the errors for jwork user package
func (c create) Err() error {
	return c.err
}

// Resp returns response for jwork user package
func (c create) Resp() io.Reader {
	return c.resp
}

// NewCreate returns User create worker
func NewCreate(dt user.Create) jwork.Worker {
	return &create{
		apiDt: dt,
	}
}

// Verification email sender
func (c create) SendVerificationMail() {
	if c.hasErr() {
		return
	}

	from := viper.GetString("mail_service.mail_address")
	pass := viper.GetString("mail_service.mail_password")
	to := c.apiDt.Email
	body := "Daakghar Verification link: http://" + viper.GetString("basic.host") + ":" + viper.GetString("basic.port") + "/api/accountverify/" + c.id.ID

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: DaakGhar Service Account Verification\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:25",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		c.err = errors.Wrap(err, " jwork.customer.create.sendVerificationMain, could not send mail ")
		return
	}
}
