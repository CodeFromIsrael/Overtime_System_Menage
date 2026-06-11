package repository

import (
	"database/sql"
	"overtime_system_menagement/src/models"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *User {
	return &User{db}
}

func (u User) Create(user models.Users) (uint64, error) {

	smt, erro := u.db.Prepare("insert into users (name, display_name, email, password, cpf, phone, role_id) (?,?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}

	defer smt.Close()

	insert, erro := smt.Exec(user.Name, user.DisplayName, user.Email, user.Password, user.Cpf, user.Phone, user.Permissions)

	if erro != nil {

		return 0, erro
	}

	lastidInser, erro := insert.LastInsertId()

	if erro != nil {
		return 0, erro
	}
	return uint64(lastidInser), nil
}
