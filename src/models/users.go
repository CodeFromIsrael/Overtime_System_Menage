package models

import (
	"errors"
	"overtime_system_menagement/src/segurity"
	"regexp"
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

func (user *Users) format(step string) error {

	user.Name = strings.TrimSpace(user.Name)

	user.Cpf = clearString(user.Cpf)

	user.Phone = clearString(user.Phone)

	if step == "cadastro" {
		passHash, erro := segurity.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passHash)
	}
	return nil
}

func clearString(str string) string {
	clear := regexp.MustCompile(`[^a-zA-Z0-9]`)
	newString := clear.ReplaceAllString(str, "")
	newString = strings.ToLower(newString)

	return newString
}
