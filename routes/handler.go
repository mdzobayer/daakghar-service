package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daakghar-service/api"
	"github.com/daakghar-service/apperr"
	"github.com/daakghar-service/config"
	"github.com/daakghar-service/token"
)

// reqLog keeps track of starting and finishing of a request handler
func reqLog(rt Route, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start to serve name=%s, path=%s, method=%s", rt.Name, rt.Path, rt.Method)
		f(w, r)
		log.Printf("done serve name=%s, path=%s, method=%s", rt.Name, rt.Path, rt.Method)
	}
}

// apiHandler handles request that does not require authentication
func apiHandler(jh api.JHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wr := jh(w, r)

		wr.Work()
		if wr.Err() != nil {
			api.ServeErr(wr.Err(), w, r)
			return
		}

		api.ServeJAPI(wr.Resp(), w, r)
	}
}

// parseCustUser returns a user from a request token
func parseCustUser(r *http.Request) (user string, err error) {
	tk := r.Header.Get("authorization")
	if tk == "" {
		return "", apperr.NewAuthentication("token")
	}

	usr := token.Decrypt(tk, config.Get().Token().Key)
	fmt.Println(usr, tk, config.Get().Token().Key)

	return usr.UserName, nil
}

// apiHandlerWithCustAuth handles request that require customer authentication
func apiHandlerWithCustAuth(f func(usr string) api.JHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := parseCustUser(r)

		if err != nil {
			api.ServeErr(err, w, r)
		}

		fmt.Println("\n\n user: ", user)
		ah := apiHandler(f(user))
		ah(w, r)
	}
}
