package router

import (
	"net/http"
	"overtime_system_menagement/src/router/routes"
)

func Generete() *http.ServeMux {
	r := http.NewServeMux()

	return routes.Config(r)
}
