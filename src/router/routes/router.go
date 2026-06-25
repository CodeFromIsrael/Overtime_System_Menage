package routes

import (
	"net/http"
	"overtime_system_menagement/src/container"
	"overtime_system_menagement/src/middleweres"
)

type Route struct {
	Uri                        string
	Method                     string
	Function                   func(http.ResponseWriter, *http.Request)
	RequiredAutentication      bool
	RequiredAdminAutentication bool
}

func Config(r *http.ServeMux, deps container.Dependences) *http.ServeMux {

	route := UsersRoutes(deps.UserController)

	route = append(route, CompanyRoutes(deps.CompanyController)...)

	route = append(route, ContractRoutes(deps.ContractController)...)

	route = append(route, allocationsRoutes(deps.AllocationController)...)

	route = append(route, routesOvertimeRecord(deps.OvertimeRecord)...)

	for _, rota := range route {

		handler := rota.Function

		if rota.RequiredAdminAutentication {
			handler = middleweres.AutenticationByAdmin(*deps.UserService)(handler)
		}

		if rota.RequiredAutentication {
			handler = middleweres.Autentication(handler)
		}

		handler = middleweres.Logger(handler)

		pattern := rota.Method + " " + rota.Uri

		r.HandleFunc(pattern, handler)
	}

	return r
}
