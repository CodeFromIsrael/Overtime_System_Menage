package bootstrap

import (
	"database/sql"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/service"
)

func buildCompanyModule(db *sql.DB) *controllers.CompanyController {

	companyRepository := repository.NewRepositoryCompany(db)

	companyService := service.NewCompanyService(companyRepository)

	companyController := controllers.NewControlerCompany(companyService)

	return companyController
}
