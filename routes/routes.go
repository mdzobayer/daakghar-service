package routes

import "net/http"

// Route holds route information
type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}
