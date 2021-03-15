package api

import (
	"encoding/json"
	"net/http"

	"github.com/daakghar-service/data/api/filters"
	"github.com/daakghar-service/data/api/user"
	duser "github.com/daakghar-service/data/api/user"
	"github.com/daakghar-service/jwork"
	jusr "github.com/daakghar-service/jwork/user"
	"github.com/pkg/errors"
)

// UsrCreate handles
func UsrCreate(w http.ResponseWriter, r *http.Request) jwork.Worker {

	usr := user.Create{}
	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		return jwork.NewErr(errors.Wrap(err, "api.UserCreate, could not decode user create request"))
	}

	return jusr.NewCreate(usr)
}

// AdmUsrGet handles get user for admin user
func AdmUsrGet(user string) JHandler {
	return func(w http.ResponseWriter, r *http.Request) jwork.Worker {
		id := filters.ID{}
		if err := id.Parse(r); err != nil {
			return jwork.NewErr(errors.Wrap(err, "api.AdmUsrGet, parse id param"))
		}

		return jusr.NewAdmGet(user, id)
	}
}

// UsrGet handles get for Customer user
func UsrGet(user string) JHandler {
	return func(w http.ResponseWriter, r *http.Request) jwork.Worker {
		id := filters.ID{}
		if err := id.Parse(r); err != nil {
			return jwork.NewErr(errors.Wrap(err, "api.UsrGet, parse id param"))
		}

		return jusr.NewUsrRead(user, id)
	}
}

// UsrUpdate handles update user information
func UsrUpdate(user string) JHandler {
	return func(w http.ResponseWriter, r *http.Request) jwork.Worker {
		id := filters.ID{}
		usr := duser.Update{}
		err := json.NewDecoder(r.Body).Decode(&usr)

		if err != nil {
			return jwork.NewErr(errors.Wrap(err, "api.UserUpdate, could not decode user update request"))
		}

		if err := id.Parse(r); err != nil {
			return jwork.NewErr(errors.Wrap(err, "api.UserPubGet, parseid param"))
		}

		return jusr.NewUpdate(id, usr)
	}
}

// UsrToken handles to generate user token
func UsrToken(w http.ResponseWriter, r *http.Request) jwork.Worker {
	tk := user.TokenReq{}
	err := json.NewDecoder(r.Body).Decode(&tk)

	if err != nil {
		return jwork.NewErr(errors.Wrap(err, "api.UsrToken, could not decode user token request"))
	}

	return jusr.NewToken(tk)
}
