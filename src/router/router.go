package router

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/router/routes"
	"overtime_system_menagement/src/service"
)

func Generete(uc *controllers.UserController, us service.UsersServices, cc *controllers.CompanyController) *http.ServeMux {
	r := http.NewServeMux()

	return routes.Config(r, uc, us, cc)
}
