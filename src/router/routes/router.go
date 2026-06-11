package routes

import (
	"net/http"
	"overtime_system_menagement/src/milddleweres"
)

type Route struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequiredAutentication bool
}

func Config(r *http.ServeMux) *http.ServeMux {

	route := routeUsers

	for _, rota := range route {

		handler := rota.Function

		handler = milddleweres.Logger(handler)

		pattern := rota.Method + " " + rota.Uri

		r.HandleFunc(pattern, handler)
	}

	return r
}
