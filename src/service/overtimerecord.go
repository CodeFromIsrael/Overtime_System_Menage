package service

import (
	"errors"
	"overtime_system_menagement/src/dto/responses"
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
	"strings"
	"time"
)

type OvertimeRecordService struct {
	repository *repository.OvertimeRecord
}

func NewServiceOvertimeRecord(repositoryOvertime *repository.OvertimeRecord) *OvertimeRecordService {
	return &OvertimeRecordService{repositoryOvertime}
}

func (o *OvertimeRecordService) CreateOvertimeRecord(ovt models.OvertimeRecord) (models.OvertimeRecord, error) {

	if err := ovt.Prepare(); err != nil {
		return models.OvertimeRecord{}, err
	}

	typeOvertime, err := o.checkTypeOvertime(ovt.WorkDate)

	if err != nil {
		return models.OvertimeRecord{}, err
	}

	ovt.OvertimeTypesId = typeOvertime

	ovt.TotalHours, err = o.calculateOvertimeHours(ovt.StartTime, ovt.EndTime)

	if err != nil {
		return models.OvertimeRecord{}, err
	}

	ovt.NigthHours, err = o.calculeteNigthHours(ovt.StartTime, ovt.EndTime)

	if err != nil {
		return models.OvertimeRecord{}, err
	}

	ovt.Id, err = o.repository.CreateOvertimeRecord(ovt)

	if err != nil {
		return models.OvertimeRecord{}, err
	}

	return ovt, nil

}

func (o *OvertimeRecordService) checkTypeOvertime(date time.Time) (uint64, error) {

	var typeId uint64

	holiday, err := o.repository.CheckIfHoliday(date)

	if err != nil {
		return 0, err
	}

	if date.Weekday() == time.Sunday {
		typeId = 2
		return typeId, nil
	}

	y1, m1, d1 := date.Date()

	y2, m2, d2 := holiday.Date.Date()

	if y1 == y2 && m1 == m2 && d1 == d2 {
		return 2, nil
	}

	typeId = 1

	return typeId, nil
}

func (o *OvertimeRecordService) calculateOvertimeHours(startTime, endTime time.Time) (float64, error) {

	duration := endTime.Sub(startTime)

	totalMinutes := duration.Minutes()

	convertedValue := totalMinutes / 60.0

	return convertedValue, nil

}

func (o *OvertimeRecordService) calculeteNigthHours(startTime, endTime time.Time) (float64, error) {

	nightStart := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 22, 0, 0, 0, startTime.Location())

	nightEnd := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 23, 59, 59, 0, startTime.Location())

	actualNightStart := startTime

	if nightStart.After(actualNightStart) {

		actualNightStart = nightStart
	}

	actualNightEnd := endTime

	if nightEnd.Before(actualNightEnd) {

		actualNightEnd = nightEnd
	}

	if actualNightStart.Before(actualNightEnd) {

		duration := actualNightEnd.Sub(actualNightStart)

		return duration.Hours(), nil
	}

	return 0, nil
}

func (o *OvertimeRecordService) ReturnOvertimeEmployee(nameEmployee string) ([]responses.OvertimeEmployee, error) {

	nameEmployee = strings.ToLower(nameEmployee)

	overtimeEmployee, err := o.repository.ReturnOvertimeEmployee(nameEmployee)

	if err != nil {
		return []responses.OvertimeEmployee{}, err
	}

	return overtimeEmployee, nil
}

func (o *OvertimeRecordService) ReturnOvertimebyId(idOvertime uint64) (responses.OvertimeEmployee, error) {

	if idOvertime == 0 {
		return responses.OvertimeEmployee{}, errors.New("id inválido")
	}

	overtimereimbursed, err := o.repository.ReturnOvertimeById(idOvertime)

	if err != nil {
		return responses.OvertimeEmployee{}, err
	}

	return overtimereimbursed, nil
}
