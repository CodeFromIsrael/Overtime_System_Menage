package controllers

import (
	"fmt"
	"net/http"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

type CompanyController struct {
	CompanyService *service.CompanyService
}

func NewControlerCompany(companySevice *service.CompanyService) *CompanyController {
	return &CompanyController{companySevice}
}

func (c *CompanyController) CreateCompany(w http.ResponseWriter, r *http.Request) {

	companyType, err := readParameter(r, "type")

	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	boryRequest, err := readerOfRequestBody(r)

	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	companyCreatedDate, err := ConverterJsonToStruct[models.Company](boryRequest)

	if err != nil {

		fmt.Println(companyCreatedDate)
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	company, err := c.CompanyService.CreateCompany(companyCreatedDate, companyType)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusCreated, company)
}
