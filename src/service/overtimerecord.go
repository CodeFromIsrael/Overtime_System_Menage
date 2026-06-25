package service

import (
	"overtime_system_menagement/src/models"
	"overtime_system_menagement/src/repository"
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

	if date.Weekday() == time.Monday || date.Equal(holiday.Date) {
		typeId = 1
		return typeId, nil
	}

	typeId = 2
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
