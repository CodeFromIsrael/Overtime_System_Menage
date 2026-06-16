package repository

import (
	"database/sql"
	"overtime_system_menagement/src/dto/responses"
	"overtime_system_menagement/src/models"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *User {
	return &User{db}
}

func (u *User) Create(user models.Users) (uint64, error) {

	smt, erro := u.db.Prepare("insert into users (name, display_name, email, password, cpf, phone, role_id) values (?,?,?,?,?,?,?)")

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

func (u *User) SearchByEmail(email string) (models.Users, error) {

	query, erro := u.db.Query("select id,password from users where email = ? ", email)

	if erro != nil {

		return models.Users{}, erro
	}

	defer query.Close()

	var user models.Users

	if query.Next() {

		if erro = query.Scan(&user.Id, &user.Password); erro != nil {

			return models.Users{}, erro
		}
	}
	return user, nil
}

func (u *User) SearchByAdmin(userid uint64) (models.Users, error) {

	lines, erro := u.db.Query("select id,name,display_name,email,phone,cpf,permissions users where id = ?", userid)

	if erro != nil {
		return models.Users{}, erro
	}

	defer lines.Close()

	var user models.Users

	if lines.Next() {
		if erro = lines.Scan(
			&user.Id,
			&user.Name,
			&user.DisplayName,
			&user.Email,
			&user.Phone,
			&user.Cpf,
			&user.Permissions,
		); erro != nil {
			return models.Users{}, erro
		}

	}

	return user, nil

}

func (u *User) CheckAdmimRole(userId uint64) (responses.UserAndRole, error) {

	query, err := u.db.Query(`select

		us.id,
		us.name,
		us.email,
		us.phone,
		us.cpf,


		rr.id,
		rr.name

		from users us
		left join roles rr 
		on rr.id = us.role_id where us.id = ?

	`, userId)

	if err != nil {
		return responses.UserAndRole{}, err
	}

	var user responses.UserAndRole

	if query.Next() {

		if err = query.Scan(
			&user.User.Id,
			&user.User.Name,
			&user.User.Email,
			&user.User.Phone,
			&user.User.Cpf,
			&user.Role.Id,
			&user.Role.Name,
		); err != nil {
			return responses.UserAndRole{}, err
		}
	}

	return user, nil
}
