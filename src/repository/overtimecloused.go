package repository

import (
	"database/sql"
	"overtime_system_menagement/src/models"
)

type OvertimeClousedRepository struct {
	db *sql.DB
}

func NewRepositoryClousedOvertime(db *sql.DB) *OvertimeClousedRepository {
	return &OvertimeClousedRepository{db}
}

func (oc *OvertimeClousedRepository) MonthdyMonthByTechLeader(clused models.OvertimeCloused) (uint64, error) {

	smt, err := oc.db.Prepare("insert into overtime_closures (company_contract_id,cloused_month,state,cloused_by,cloused_at) values (?,?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer smt.Close()

	insert, err := smt.Exec(clused.ContractCompanyId, clused.PeriodCloused, clused.State, clused.ClousedBy, clused.ClousedAt)

	if err != nil {
		return 0, err
	}

	lastidinser, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastidinser), err
}
