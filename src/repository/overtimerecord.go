package repository

import (
	"database/sql"
	"overtime_system_menagement/src/dto/responses"
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

func (o *OvertimeRecord) ReturnOvertimeEmployee(name string) ([]responses.OvertimeEmployee, error) {

	query, err := o.db.Query(`
	  select 

	o.id AS overtime_records_id,
    o.work_date As overtime_records_work_date,
    o.start_time AS overtime_records_start_time,
    o.end_time AS overtime_records_end_time,
    o.overtime_type_id AS overtime_records_type_overtime_id,
    o.total_hours AS overtime_records_total_hours,
    o.night_hours AS overtime_night_hours,
    
    u.id AS users_id,
    u.name AS users_name
    
    from overtime_records o
    inner join allocations al
    
    on o.allocation_id = al.id 
    inner join contracts_employee ce
    
    on al.employee_contract_id = ce.id
    
    inner join users u
    on ce.user_id = u.id 
    
    where u.name = ? order by o.work_date desc;
	`, name)

	if err != nil {
		return []responses.OvertimeEmployee{}, err
	}

	defer query.Close()

	var overtimeEmployee []responses.OvertimeEmployee

	for query.Next() {

		var oE responses.OvertimeEmployee

		if err = query.Scan(
			&oE.Overtime.Id,
			&oE.Overtime.WorkDate,
			&oE.TypeStartTimeReturned,
			&oE.TypeEndtimeReturned,
			&oE.Overtime.OvertimeTypesId,
			&oE.Overtime.TotalHours,
			&oE.Overtime.NigthHours,
			&oE.Employee.Id,
			&oE.Employee.Name,
		); err != nil {
			return []responses.OvertimeEmployee{}, err
		}

		overtimeEmployee = append(overtimeEmployee, oE)
	}

	return overtimeEmployee, nil
}

func (o *OvertimeRecord) ReturnOvertimeById(id uint64) (responses.OvertimeEmployee, error) {

	query, err := o.db.Query(`
	  select 

	o.id AS overtime_records_id,
    o.work_date As overtime_records_work_date,
    o.start_time AS overtime_records_start_time,
    o.end_time AS overtime_records_end_time,
    o.overtime_type_id AS overtime_records_type_overtime_id,
    o.total_hours AS overtime_records_total_hours,
    o.night_hours AS overtime_night_hours,
    
    u.id AS users_id,
    u.name AS users_name
    
    from overtime_records o
    inner join allocations al
    
    on o.allocation_id = al.id 
    inner join contracts_employee ce
    
    on al.employee_contract_id = ce.id
    
    inner join users u
    on ce.user_id = u.id 
    
    where o.id = ? 
	`, id)

	if err != nil {
		return responses.OvertimeEmployee{}, err
	}

	defer query.Close()

	var overtimeEmployee responses.OvertimeEmployee

	for query.Next() {

		if err = query.Scan(
			&overtimeEmployee.Overtime.Id,
			&overtimeEmployee.Overtime.WorkDate,
			&overtimeEmployee.TypeStartTimeReturned,
			&overtimeEmployee.TypeEndtimeReturned,
			&overtimeEmployee.Overtime.OvertimeTypesId,
			&overtimeEmployee.Overtime.TotalHours,
			&overtimeEmployee.Overtime.NigthHours,
			&overtimeEmployee.Employee.Id,
			&overtimeEmployee.Employee.Name,
		); err != nil {
			return responses.OvertimeEmployee{}, err
		}

	}

	return overtimeEmployee, nil
}
