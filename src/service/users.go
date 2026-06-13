package service

import (
	"fmt"
	"overtime_system_menagement/src/autentication"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/segurity"
	"strconv"
)

type UsersServices struct {
	userRepository *repository.User
}

func NewUserService(userRepository *repository.User) *UsersServices {
	return &UsersServices{userRepository}
}

func (s *UsersServices) CreateUser(user models.Users) (models.Users, error) {

	if erro := user.Prepare("cadastro"); erro != nil {
		fmt.Println("não mandei o cadastro ")
		return user, erro
	}

	var err error

	user.Id, err = s.userRepository.Create(user)

	if err != nil {
		return user, err

	}

	return user, nil
}

func (us *UsersServices) Login(user models.Users) (models.AuthenticationData, error) {

	userInDatabase, err := us.userRepository.SearchByEmail(user.Email)

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
