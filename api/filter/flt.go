package filter

import (
	"net/http"

	"github.com/daakghar-service/apperr"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

// ID handles id filter
type ID struct {
	Val string
}

func (id *ID) valid() error {
	if !bson.IsObjectIdHex(id.Val) {
		return apperr.NewValidation("id", "invalid id", "pass proper bson id")
	}

	return nil
}

// Parse parse id from get params
func (id *ID) Parse(r *http.Request) error {
	id.Val = mux.Vars(r)["id"]

	if err := id.valid(); err != nil {
		return errors.Wrap(err, "data.api.filters.ID.Parse")
	}

	return nil
}
