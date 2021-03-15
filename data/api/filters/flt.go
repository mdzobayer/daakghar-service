package filters

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/daakghar-service/apperr"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// ID handles id filter
type ID struct {
	Val string
}

func (id ID) valid() error {
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

// Skip handles skip filter
type Skip struct {
	Val int
}

// Parse parse skip from get params
func (s *Skip) Parse(r *http.Request) error {
	v, err := strconv.Atoi(mux.Vars(r)["skip"])

	if err != nil {
		return errors.Wrapf(
			apperr.NewValidation("skip", "invalid number", ""),
			"data.api.filters.Skip.Parse, %v",
			err,
		)
	}

	s.Val = v

	return nil
}

// Limit handles limit filter
type Limit struct {
	Val int
}

// Parse parse skip from get params
func (l *Limit) Parse(r *http.Request) error {
	v, err := strconv.Atoi(mux.Vars(r)["limit"])

	if err != nil {
		return errors.Wrapf(
			apperr.NewValidation("limit", "invalid number", ""),
			"data.api.filters.limit.Parse, %v",
			err,
		)
	}

	l.Val = v

	return nil
}

// Sort handles sort filter
type Sort struct {
	Field string
	Dir   string
}

// Parse parse sort from get params
func (s *Sort) Parse(r *http.Request) error {
	s.Field = mux.Vars(r)["sort"]
	s.Dir = mux.Vars(r)["sort-dir"]

	return nil
}

// Search handles search filter
type Search struct {
	Val string
}

// Parse parse search from get params
func (s *Search) Parse(r *http.Request) error {
	s.Val = mux.Vars(r)["search"]
	return nil
}

// List handles list filter
type List struct {
	Skip   Skip
	Limit  Limit
	Sort   Sort
	Search Search
}

// Parse parse list params from get params
func (l *List) Parse(r *http.Request) error {
	if err := l.Skip.Parse(r); err != nil {
		return errors.Wrap(err, "data.api.filters.List, could not parse skip")
	}

	if err := l.Limit.Parse(r); err != nil {
		return errors.Wrap(err, "data.api.filters.List, could not parse limit")
	}

	if err := l.Sort.Parse(r); err != nil {
		return errors.Wrap(err, "data.api.filters.List, could not parse sort")
	}

	if err := l.Search.Parse(r); err != nil {
		return errors.Wrap(err, "data.api.filters.List, could not parse search")
	}

	return nil
}
