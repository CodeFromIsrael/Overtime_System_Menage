package controllers

import (
	"net/http"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

type ContractController struct {
	contractService *service.ContractService
}

func NewControlerContract(serviceContract *service.ContractService) *ContractController {
	return &ContractController{serviceContract}
}

func (cc *ContractController) CreateContractEmployee(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := readerOfRequestBody(r)

	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	dateContractEmployee, err := ConverterJsonToStruct[models.ContractEmployee](bodyRequest)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	contractCreadted, err := cc.contractService.CreateContractEmployee(dateContractEmployee)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusCreated, contractCreadted)
}

func (cc *ContractController) CreateContractCompany(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := readerOfRequestBody(r)

	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	dateContractCompany, err := ConverterJsonToStruct[models.ContractCompany](bodyRequest)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	contractCreadted, err := cc.contractService.CreateContractCompany(dateContractCompany)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusCreated, contractCreadted)
}
