package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/daakghar-service/apperr"
	"github.com/pkg/errors"
)

// ServeErr serve err
func ServeErr(serr error, w http.ResponseWriter, r *http.Request) {

	log.Println("ServeErr", serr)
	err := errors.Cause(serr)

	w.WriteHeader(statusCode(err))
	err = json.NewEncoder(w).Encode(struct {
		Msg string `json:"error"`
	}{Msg: errMsg(err)},
	)

}

func errMsg(err error) string {
	switch err.(type) {
	case apperr.StatusCoder:
		return err.Error()
	default:
		return "Internal server error"
	}
}

func statusCode(err error) int {
	switch err.(type) {
	case apperr.StatusCoder:
		s := err.(apperr.StatusCoder)
		return s.StatusCode()
	default:
		return http.StatusInternalServerError
	}
}

// ServeJAPI serve json data
func ServeJAPI(read io.Reader, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_, err := io.Copy(w, read)
	if err != nil {
		log.Println("error to serve json", err)
	}

}
