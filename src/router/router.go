package router

import (
	"webapp_go/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configured(r)
}
