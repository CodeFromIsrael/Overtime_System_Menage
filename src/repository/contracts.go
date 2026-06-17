package repository

import (
	"database/sql"
	"overtime_system_menagement/src/models"
)

type Contract struct {
	db *sql.DB
}

func NewRepositoryContract(db *sql.DB) *Contract {
	return &Contract{db}
}

func (ce *Contract) CreateContract(contract models.ContractEmployee) (uint64, error) {

	smt, err := ce.db.Prepare("insert into contracts_employee (admission_date,termination_date,service_company_id,user_id,status,position) values (?,?,?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer smt.Close()

	insert, err := smt.Exec(contract.AdmissionDate, contract.TerminationDate, contract.ServiceCompanyId, contract.UserId, contract.Status, contract.Position)

	if err != nil {
		return 0, err
	}
	lastidinser, erro := insert.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastidinser), erro

}

func (ce *Contract) CreateContractCompany(contract models.ContractCompany) (uint64, error) {

	smt, err := ce.db.Prepare("insert into contracts_company (service_company_id,client_company_id,start_date,end_date,status) values (?,?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer smt.Close()

	insert, err := smt.Exec(contract.ServiceCompanyId, contract.ClientCompanyId, contract.StartTime, contract.EndTime, contract.State)

	if err != nil {
		return 0, err
	}

	lastidinser, erro := insert.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastidinser), erro

}
