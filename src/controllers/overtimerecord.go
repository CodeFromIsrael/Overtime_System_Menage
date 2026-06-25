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

	overtimeCreated, err := o.serviceOvertimeRecord.CreateOvertimeRecord(dataOvertime)

	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	response.Json(w, http.StatusCreated, overtimeCreated)
}
