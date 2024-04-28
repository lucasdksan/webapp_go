package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI                     string
	Method                  string
	Functionality           func(http.ResponseWriter, *http.Request)
	Requires_authentication bool
}

func Configured(r *mux.Router) *mux.Router {
	routes := Login_route
	routes = append(routes, User_routes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Functionality).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
