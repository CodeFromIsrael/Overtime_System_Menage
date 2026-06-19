package bootstrap

import (
	"database/sql"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/service"
)

func buildContractModule(db *sql.DB) *controllers.ContractController {

	contractRepository := repository.NewRepositoryContract(db)

	contractService := service.NewContractService(contractRepository)

	contractController := controllers.NewControlerContract(contractService)

	return contractController
}
