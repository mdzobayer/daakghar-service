package routes

import (
	"github.com/daakghar-service/api"
	"github.com/gorilla/mux"
)

// Get returns all routes
func Get() *mux.Router {
	r := mux.NewRouter()

	rts := []Route{}
	user(&rts)

	for _, v := range rts {
		r.HandleFunc(v.Path, reqLog(v, v.Handler)).Name(v.Name).Methods(v.Method)
	}

	return r
}

// user creates all user api routes
func user(r *[]Route) {
	*r = append(*r, Route{
		Name:    "Create User",
		Path:    "/api/user",
		Method:  "POST",
		Handler: apiHandler(api.UsrCreate),
	})

	*r = append(*r, Route{
		Name:    "Get User login token",
		Path:    "/api/token",
		Method:  "POST",
		Handler: apiHandler(api.UsrToken),
	})

	*r = append(*r, Route{
		Name:    "Get User",
		Path:    "/api/user/{id}",
		Method:  "GET",
		Handler: apiHandlerWithCustAuth(api.UsrGet),
	})

	*r = append(*r, Route{
		Name:    "Get user for admin user",
		Path:    "/api/admin/user/{id}",
		Method:  "GET",
		Handler: apiHandlerWithCustAuth(api.AdmUsrGet),
	})

	*r = append(*r, Route{
		Name:    "Update user info for user",
		Path:    "/api/user/update/{id}",
		Method:  "POST",
		Handler: apiHandlerWithCustAuth(api.UsrUpdate),
	})
}
