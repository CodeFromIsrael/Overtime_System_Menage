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
