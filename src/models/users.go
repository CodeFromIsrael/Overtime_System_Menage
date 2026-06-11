package models

import (
	"errors"
	"overtime_system_menagement/src/segurity"
	"strings"
)

type Users struct {
	Id          uint64  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Email       string  `json:"email,omitempty"`
	Password    string  `json:"password,omitempty"`
	Cpf         string  `json:"cpf,omitempty"`
	Phone       string  `json:"phone,omitempty"`
	Permissions uint64  `json:"role_id,omitempty"`
}

func (u *Users) Prepare(step string) error {
	if erro := u.validate(); erro != nil {
		return erro
	}
	if erro := u.format(step); erro != nil {
		return erro
	}

	return nil
}

func (u *Users) validate() error {

	if u.Name == "" {
		return errors.New("campo nome e obrigatório")
	}

	u.displayName()

	if u.Email == "" {
		return errors.New("campo email vazio")
	}

	if u.Password == "" {
		return errors.New("campo senha vazio")
	}

	if u.Cpf == "" {
		return errors.New("campo cpf vazio")
	}

	if u.Phone == "" {
		return errors.New("campo telefone vazio")
	}

	return nil

}

func (u *Users) displayName() string {

	if u.DisplayName == nil {
		return ""
	}

	return *u.DisplayName
}

func (User *Users) format(step string) error {

	User.Name = strings.TrimSpace(User.Name)

	//User.Name = clearString(User.Name)
	//User.Number = clearString(User.Number)
	if step == "cadastro" {
		passHash, erro := segurity.Hash(User.Password)
		if erro != nil {
			return erro
		}
		User.Password = string(passHash)
	}
	return nil
}
