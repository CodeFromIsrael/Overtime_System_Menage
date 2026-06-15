package service

import (
	"errors"
	"fmt"
	"overtime_system_menagement/src/autentication"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
	"overtime_system_menagement/src/response/responses"
	"overtime_system_menagement/src/segurity"
	"strconv"
)

type UsersServices struct {
	userRepository *repository.User
}

func NewUserService(userRepository *repository.User) *UsersServices {
	return &UsersServices{userRepository}
}

func (s *UsersServices) CreateUser(user models.Users) (responses.CreateUserResponse, error) {

	if erro := user.Prepare("cadastro"); erro != nil {
		fmt.Println("não mandei o cadastro ")
		return responses.CreateUserResponse{}, erro
	}

	var err error

	user.Id, err = s.userRepository.Create(user)

	if err != nil {
		return responses.CreateUserResponse{}, err

	}

	var dateUserRetorned responses.CreateUserResponse

	dateUserRetorned.Id = user.Id
	dateUserRetorned.Name = user.Name
	dateUserRetorned.Email = user.Email
	dateUserRetorned.Phone = user.Phone

	return dateUserRetorned, nil
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

func (us *UsersServices) CheckPermissionByAdmin(userid uint64) error {

	userInRequest, err := us.userRepository.CheckAdmimRole(userid)

	if err != nil {
		return err
	}

	if userInRequest.Role.Name != "Super_Administrator" {
		return errors.New("Usuário sem Permisão")
	}

	return nil
}
