package controllers

import (
	"fmt"
	"net/http"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	bodyrequest, erro := readerOfRequestBody(r)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userCreated, erro := ConverterJsonToStruct[models.Users](bodyrequest)

	service := service.NewUserService()

	create, erro := service.CreateUser(userCreated)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("problema na função no service da função ")
		return
	}

	response.Json(w, http.StatusCreated, create)
}

func Login(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := readerOfRequestBody(r)

	if err != nil {

		response.Erro(w, http.StatusBadRequest, err)

		return
	}

	user, err := ConverterJsonToStruct[models.Users](bodyRequest)

	if err != nil {

		response.Erro(w, http.StatusConflict, err)

		return
	}

	service := service.NewUserService()

	tokenUser, err := service.Login(user)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusOK, tokenUser)
}
