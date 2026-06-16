package bootstrap

import (
	"database/sql"
	"net/http"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/router"
	"overtime_system_menagement/src/service"
)

func Initialize(Dd *sql.DB) *http.ServeMux {

	userRepo := repository.NewRepositoryUser(Dd)

	userService := service.NewUserService(userRepo)

	userController := controllers.NewControlerUser(userService)

	companyRepository := repository.NewRepositoryCompany(Dd)

	companyService := service.NewCompanyService(companyRepository)

	companyController := controllers.NewControlerCompany(companyService)

	r := router.Generete(userController, *userService, companyController)

	return r
}
