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
	{
		URI:                     "/publications/{id}/like",
		Method:                  http.MethodPost,
		Functionality:           controllers.Like_Publication,
		Requires_authentication: true,
	},
	{
		URI:                     "/publications/{id}/dislike",
		Method:                  http.MethodPost,
		Functionality:           controllers.Dislike_Publication,
		Requires_authentication: true,
	},
	{
		URI:                     "/publications/{id}/update",
		Method:                  http.MethodGet,
		Functionality:           controllers.Load_Update_Post_Screen,
		Requires_authentication: true,
	},
	{
		URI:                     "/publications/{id}",
		Method:                  http.MethodPut,
		Functionality:           controllers.Update_Publication,
		Requires_authentication: true,
	},
	{
		URI:                     "/publications/{id}",
		Method:                  http.MethodDelete,
		Functionality:           controllers.Delete_Publication,
		Requires_authentication: true,
	},
}
