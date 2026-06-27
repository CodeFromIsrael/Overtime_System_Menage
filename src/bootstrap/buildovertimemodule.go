package bootstrap

import (
	"database/sql"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/service"
)

func buildOvertimeModule(db *sql.DB) (*controllers.OvertimeRecordController, *service.OvertimeRecordService) {

	overtimeRepository := repository.NewRepositoryOvertimeRecord(db)

	overtimeService := service.NewServiceOvertimeRecord(overtimeRepository)

	overtimeController := controllers.NewControllerOvertimeRecord(overtimeService)

	return overtimeController, overtimeService
}
