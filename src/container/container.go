package container

import (
	"overtime_system_menagement/src/controllers"
	"overtime_system_menagement/src/service"
)

type Dependences struct {
	UserController       *controllers.UserController
	UserService          *service.UsersServices
	CompanyController    *controllers.CompanyController
	ContractController   *controllers.ContractController
	AllocationController *controllers.Allocations
}
