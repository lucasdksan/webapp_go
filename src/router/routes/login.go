package routes

import (
	"net/http"
	"webapp_go/src/controllers"
)

var Login_route = []Routes{
	{
		URI:                     "/",
		Method:                  http.MethodGet,
		Functionality:           controllers.Load_Login_Screen,
		Requires_authentication: false,
	},
	{
		URI:                     "/login",
		Method:                  http.MethodGet,
		Functionality:           controllers.Load_Login_Screen,
		Requires_authentication: false,
	},
	{
		URI:                     "/login",
		Method:                  http.MethodPost,
		Functionality:           controllers.Handle_Login,
		Requires_authentication: false,
	},
}
