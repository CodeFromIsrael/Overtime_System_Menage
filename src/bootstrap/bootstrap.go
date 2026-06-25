package bootstrap

import (
	"database/sql"
	"net/http"
	"overtime_system_menagement/src/container"
	"overtime_system_menagement/src/router"
)

func Initialize(Dd *sql.DB) *http.ServeMux {

	userController, userService := builUserModule(Dd)

	companyController := buildCompanyModule(Dd)

	contractController := buildContractModule(Dd)

	allocationControllers := buildAllocationModule(Dd)

	overtimeControllers := buildOvertimeModule(Dd)

	deps := container.Dependences{
		UserController:       userController,
		UserService:          userService,
		ContractController:   contractController,
		CompanyController:    companyController,
		AllocationController: allocationControllers,
		OvertimeRecord:       overtimeControllers,
	}

	r := router.Generete(deps)

	return r
}
