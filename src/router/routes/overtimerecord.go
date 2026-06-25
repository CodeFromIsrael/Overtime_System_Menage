package routes

import (
	"net/http"
	"overtime_system_menagement/src/controllers"
)

func routesOvertimeRecord(co *controllers.OvertimeRecordController) []Route {

	return []Route{
		{
			Uri:                   "/overtime",
			Method:                http.MethodPost,
			Function:              co.CreateOvertimeRecord,
			RequiredAutentication: true,
		},
	}
}
