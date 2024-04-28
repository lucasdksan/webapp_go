package routes

import (
	"net/http"
	"webapp_go/src/controllers"
)

var User_routes = []Routes{
	{
		URI:                     "/user-registration",
		Method:                  http.MethodGet,
		Functionality:           controllers.Load_Registration_Screen,
		Requires_authentication: false,
	},
	{
		URI:                     "/users",
		Method:                  http.MethodPost,
		Functionality:           controllers.Create_User,
		Requires_authentication: false,
	},
}
