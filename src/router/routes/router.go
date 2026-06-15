package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/middleweres"
	"overtime_system_menagement/src/service"
)

type Route struct {
	Uri                        string
	Method                     string
	Function                   func(http.ResponseWriter, *http.Request)
	RequiredAutentication      bool
	RequiredAdminAutentication bool
}

func Config(r *http.ServeMux, uc *controllers.UserController, us service.UsersServices) *http.ServeMux {

	route := UsersRoutes(uc)

	for _, rota := range route {

		handler := rota.Function

		if rota.RequiredAutentication {
			handler = middleweres.Autentication(handler)
		}

		if rota.RequiredAdminAutentication {
			handler = middleweres.AutenticationByAdmin(us)(handler)
		}

		handler = middleweres.Logger(handler)

		pattern := rota.Method + " " + rota.Uri

		r.HandleFunc(pattern, handler)
	}

	return r
}
