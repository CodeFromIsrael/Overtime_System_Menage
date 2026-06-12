package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
)

var routeUsers = []Route{
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUsers,
		RequiredAutentication: false,
	},
	{
		Uri:                   "/login",
		Method:                http.MethodPost,
		Function:              controllers.Login,
		RequiredAutentication: false,
	},
}
