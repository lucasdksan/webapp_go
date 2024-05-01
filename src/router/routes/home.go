package routes

import (
	"net/http"
	"webapp_go/src/controllers"
)

var Home_page_route = Routes{
	URI:                     "/home",
	Method:                  http.MethodGet,
	Functionality:           controllers.Load_Home_Screen,
	Requires_authentication: true,
}
