package bootstrap

import (
	"database/sql"
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/service"
)

func builUserModule(db *sql.DB) (*controllers.UserController, *service.UsersServices) {

	userRepo := repository.NewRepositoryUser(db)

	userService := service.NewUserService(userRepo)

	userController := controllers.NewControlerUser(userService)

	return userController, userService
}
