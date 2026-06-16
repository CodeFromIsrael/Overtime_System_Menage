package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
)

func CompanyRoutes(cc *controllers.CompanyController) []Route {
	return []Route{
		{
			Uri:                        "/company",
			Method:                     http.MethodPost,
			Function:                   cc.CreateCompany,
			RequiredAutentication:      true,
			RequiredAdminAutentication: true,
		},
	}
}
