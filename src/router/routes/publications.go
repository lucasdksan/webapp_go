package routes

import (
	"net/http"
	"webapp_go/src/controllers"
)

var Publications_route = []Routes{
	{
		URI:                     "/publications",
		Method:                  http.MethodPost,
		Functionality:           controllers.Create_Publication,
		Requires_authentication: true,
	},
}