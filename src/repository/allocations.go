package repository

import (
	"database/sql"
	"overtime_system_menagement/src/models"
)

type Allocations struct {
	db *sql.DB
}

func NewRepositoryAllocations(db *sql.DB) *Allocations {
	return &Allocations{db}
}

func (a *Allocations) Create(allocation models.Allocations) (uint64, error) {

	smt, err := a.db.Prepare("insert into allocations (employee_contract_id,company_contract_id,start_date,end_date) values (?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer smt.Close()

	insert, err := smt.Exec(allocation.ContractEmployeeId, allocation.ContractCompanyId, allocation.StartDate, allocation.EndDate)

	if err != nil {
		return 0, err
	}

	lastidinser, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastidinser), err

}
