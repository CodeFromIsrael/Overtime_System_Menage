package controllers

import (
	"net/http"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

type Allocations struct {
	allocationService *service.AllocationsService
}

func NewControllerAllocations(allocationService *service.AllocationsService) *Allocations {
	return &Allocations{allocationService}
}

func (a *Allocations) CreateAllocation(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := readerOfRequestBody(r)

	if err != nil {

		response.Erro(w, http.StatusBadRequest, err)

		return
	}

	dateAllocation, err := ConverterJsonToStruct[models.Allocations](bodyRequest)

	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)

		return
	}

	allocationCreated, err := a.allocationService.CreateAllocation(dateAllocation)

	if err != nil {

		response.Erro(w, http.StatusInternalServerError, err)

		return
	}

	response.Json(w, http.StatusCreated, allocationCreated)
}
