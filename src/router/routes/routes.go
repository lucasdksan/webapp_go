package routes

import (
	"net/http"
	"webapp_go/src/middlewares"

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
	routes = append(routes, Home_page_route)
	routes = append(routes, Publications_route...)
	routes = append(routes, Logout_route)

	for _, route := range routes {
		if route.Requires_authentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Functionality))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Functionality)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
