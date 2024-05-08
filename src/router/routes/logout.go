package routes

import (
	"net/http"
	"webapp_go/src/controllers"
)

var Logout_route = Routes{
	URI:                     "/logout",
	Method:                  http.MethodGet,
	Functionality:           controllers.Handle_Logout,
	Requires_authentication: true,
}
