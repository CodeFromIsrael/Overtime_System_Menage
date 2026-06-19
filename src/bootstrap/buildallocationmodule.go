package bootstrap

import (
	"database/sql"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/service"
)

func buildAllocationModule(db *sql.DB) *controllers.Allocations {

	allocationRepository := repository.NewRepositoryAllocations(db)

	allocationSevice := service.NewAllocationService(allocationRepository)

	allocationControllers := controllers.NewControllerAllocations(allocationSevice)

	return allocationControllers
}
