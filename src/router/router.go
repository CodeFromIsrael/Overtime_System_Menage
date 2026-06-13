package router

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/router/routes"
)

func Generete(uc *controllers.UserController) *http.ServeMux {
	r := http.NewServeMux()

	return routes.Config(r, uc)
}
