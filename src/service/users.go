package service

import (
	"fmt"
	"overtime_system_menagement/src/datebase"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
)

type UsersServices struct{}

func NewUserService() *UsersServices {
	return &UsersServices{}
}

func (s *UsersServices) CreateUser(user models.Users) (models.Users, error) {

	if erro := user.Prepare("cadastro"); erro != nil {
		fmt.Println("não mandei o cadastro ")
		return user, erro
	}

	db, erro := datebase.Connection()

	if erro != nil {
		return user, erro

	}

	defer db.Close()

	repository := repository.NewRepositoryUser(db)

	user.Id, erro = repository.Create(user)

	if erro != nil {
		return user, erro

	}

	return user, nil
}
