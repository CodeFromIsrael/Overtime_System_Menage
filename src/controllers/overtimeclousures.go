package controllers

import (
	"net/http"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

type OvertimeClousedControllers struct {
	service *service.ClousuresOvertimeServices
}

func NewControllerCloused(serviceCloused *service.ClousuresOvertimeServices) *OvertimeClousedControllers {
	return &OvertimeClousedControllers{serviceCloused}
}

func (occ *OvertimeClousedControllers) CreateClosingInMonth(w http.ResponseWriter, r *http.Request) {

	userInRequest := retriveUserInToken(r)

	state, err := occ.service.CreateClosingInMonth(userInRequest)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusCreated, state)
}
