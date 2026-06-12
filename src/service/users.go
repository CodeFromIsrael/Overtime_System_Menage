package service

import (
	"fmt"
	"overtime_system_menagement/src/autentication"
	"overtime_system_menagement/src/datebase"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/segurity"
	"strconv"
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

func (us *UsersServices) Login(user models.Users) (models.AuthenticationData, error) {

	db, err := datebase.Connection()

	if err != nil {
		return models.AuthenticationData{}, err
	}

	defer db.Close()

	repository := repository.NewRepositoryUser(db)

	userInDatabase, err := repository.SearchByEmail(user.Email)

	if err != nil {
		return models.AuthenticationData{}, err
	}

	if err = segurity.CheckPass(userInDatabase.Password, user.Password); err != nil {

		return models.AuthenticationData{}, err
	}

	token, err := autentication.CreateToken(userInDatabase.Id)

	if err != nil {
		return models.AuthenticationData{}, err
	}

	userId := strconv.FormatUint(userInDatabase.Id, 10)

	return models.AuthenticationData{ID: userId, Token: token}, nil

}
