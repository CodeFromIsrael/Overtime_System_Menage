package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
)

func allocationsRoutes(alc *controllers.Allocations) []Route {
	return []Route{
		{
			Uri:                        "/allocations",
			Method:                     http.MethodPost,
			Function:                   alc.CreateAllocation,
			RequiredAutentication:      true,
			RequiredAdminAutentication: true,
		},
	}
}
