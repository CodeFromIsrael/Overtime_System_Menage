package controllers

import (
	"net/http"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/response"
	"overtime_system_menagement/src/service"
)

type OvertimeRecordController struct {
	serviceOvertimeRecord *service.OvertimeRecordService
}

func NewControllerOvertimeRecord(serviceOvertime *service.OvertimeRecordService) *OvertimeRecordController {
	return &OvertimeRecordController{serviceOvertime}
}

func (o *OvertimeRecordController) CreateOvertimeRecord(w http.ResponseWriter, r *http.Request) {

	userInRequest := retriveUserInToken(r)

	bodyRequest, err := readerOfRequestBody(r)

	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	dataOvertime, err := ConverterJsonToStruct[models.OvertimeRecord](bodyRequest)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	overtimeCreated, err := o.serviceOvertimeRecord.CreateOvertimeRecord(dataOvertime, userInRequest)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	response.Json(w, http.StatusCreated, overtimeCreated)
}

func (o *OvertimeRecordController) ReturnOvertimeEmployee(w http.ResponseWriter, r *http.Request) {

	nameEmployee, err := readParameter(r, "name")

	if err != nil {

		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	overtimeEmployee, err := o.serviceOvertimeRecord.ReturnOvertimeEmployee(nameEmployee)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.Json(w, http.StatusOK, overtimeEmployee)
}
