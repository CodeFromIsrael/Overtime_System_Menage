package repository

import (
	"database/sql"
	"errors"
	"overtime_system_menagement/src/models"
)

type Company struct {
	db *sql.DB
}

func NewRepositoryCompany(db *sql.DB) *Company {
	return &Company{db}
}

func (c *Company) CreateCompany(company models.Company, typeCompany string) (uint64, error) {

	var query string

	switch typeCompany {

	case "client":
		query = "insert into client_company (name,legal_name,tax_identifier,status) values (?,?,?,?)"

	case "service":
		query = "insert into service_company (name,legal_name,tax_identifier,status) values (?,?,?,?)"

	default:
		return 0, errors.New("opção enviada invalida")

	}

	//fmt.Println("query que esta sendo executada:", query)

	smt, err := c.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer smt.Close()

	insert, err := smt.Exec(company.Name, company.LegalName, company.TaxIndefier, company.State)

	if err != nil {
		return 0, err
	}

	lastidinser, erro := insert.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastidinser), erro
}
