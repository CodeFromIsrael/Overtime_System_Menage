package repository

import (
	"database/sql"
	"overtime_system_menagement/src/models"
	"time"
)

type OvertimeRecord struct {
	db *sql.DB
}

func NewRepositoryOvertimeRecord(db *sql.DB) *OvertimeRecord {
	return &OvertimeRecord{db}
}

func (o *OvertimeRecord) CreateOvertimeRecord(ovtr models.OvertimeRecord) (uint64, error) {

	smt, err := o.db.Prepare("insert into overtime_records (work_date,start_time,end_time,jira_task_identifier,observation,total_hours,night_hours,overtime_type_id,allocation_id) values (?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return 0, err
	}

	defer smt.Close()

	insert, err := smt.Exec(ovtr.WorkDate, ovtr.StartTime, ovtr.EndTime, ovtr.JiraTaskIndefier, ovtr.Observation, ovtr.TotalHours, ovtr.NigthHours, ovtr.OvertimeTypesId, ovtr.AllocationId)

	if err != nil {
		return 0, err
	}

	lastidInser, erro := insert.LastInsertId()

	if erro != nil {
		return 0, erro
	}
	return uint64(lastidInser), nil
}

func (o *OvertimeRecord) CheckIfHoliday(date time.Time) (models.Holiday, error) {

	query, err := o.db.Query("select holiday_date from holiday where holiday_date = ?", date)

	if err != nil {
		return models.Holiday{}, err
	}

	defer query.Close()

	var holiday models.Holiday

	if query.Next() {

		if err = query.Scan(
			&holiday.Date,
		); err != nil {
			return models.Holiday{}, err
		}
	}

	return holiday, nil
}
