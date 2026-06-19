package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
)

func ContractRoutes(cc *controllers.ContractController) []Route {
	return []Route{
		{
			Uri:                        "/contract/employee",
			Method:                     http.MethodPost,
			Function:                   cc.CreateContractEmployee,
			RequiredAutentication:      true,
			RequiredAdminAutentication: true,
		},
		{
			Uri:                        "/contract/company",
			Method:                     http.MethodPost,
			Function:                   cc.CreateContractCompany,
			RequiredAutentication:      true,
			RequiredAdminAutentication: true,
		},
	}
}
