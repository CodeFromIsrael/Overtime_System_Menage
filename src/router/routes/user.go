package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
)

func UsersRoutes(uc *controllers.UserController) []Route {
	return []Route{
		{
			Uri:                        "/users",
			Method:                     http.MethodPost,
			Function:                   uc.CreateUsers,
			RequiredAutentication:      false,
			RequiredAdminAutentication: false,
		},
		{
			Uri:                   "/login",
			Method:                http.MethodPost,
			Function:              uc.Login,
			RequiredAutentication: false,
		},
	}

}
