package router

import (
	"net/http"
	"overtime_system_menagement/src/container"
	"overtime_system_menagement/src/router/routes"
)

func Generete(depences container.Dependences) *http.ServeMux {
	r := http.NewServeMux()

	return routes.Config(r, depences)
}
