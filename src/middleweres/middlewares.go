package middleweres

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"overtime_system_menagement/src/autentication"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		nextFunc(w, r)
	}
}

func Autentication(nextFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if erro := autentication.ValidadeToken(r); erro != nil {

			fmt.Println("não valida o token")

			response.Erro(w, http.StatusUnauthorized, erro)

			return
		}

		userInToken, erro := autentication.ExtractUserId(r)

		if erro != nil {
			response.Erro(w, http.StatusBadRequest, erro)
			return
		}

		context := context.WithValue(r.Context(), "userid", userInToken)

		r = r.WithContext(context)

		nextFunc(w, r)
	}
}

func AutenticationByAdmin(userService service.UsersServices) func(http.HandlerFunc) http.HandlerFunc {

	return func(nextFunc http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			userInRequest, err := autentication.ExtractUserId(r)

			if err != nil {
				response.Erro(w, http.StatusBadRequest, err)
				return
			}

			if err = userService.CheckPermissionByAdmin(userInRequest); err != nil {
				response.Erro(w, http.StatusUnauthorized, err)
				return
			}

			nextFunc(w, r)
		}

	}
}
