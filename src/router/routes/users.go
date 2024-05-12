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
	{
		URI:                     "/users-search",
		Method:                  http.MethodGet,
		Functionality:           controllers.Load_Search_Screen,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodGet,
		Functionality:           controllers.Load_Profile_Screen,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}/unfollow",
		Method:                  http.MethodPost,
		Functionality:           controllers.Un_Follow_User,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}/follow",
		Method:                  http.MethodPost,
		Functionality:           controllers.Follow_User,
		Requires_authentication: true,
	},
}
