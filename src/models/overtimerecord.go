package models

import (
	"errors"
	"time"
)

type OvertimeRecord struct {
	Id               uint64    `json:"id,omitempty"`
	WorkDate         time.Time `json:"work_date"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	JiraTaskIndefier string    `json:"jira_task_indetifier"`
	Observation      *string   `json:"observation,omitempty"`
	TotalHours       float64   `json:"total_hours,omitempty"`
	NigthHours       float64   `json:"nigth_hours,omitempty"`
	OvertimeTypesId  uint64    `json:"overtime_type_id,omitempty"`
	AllocationId     uint64    `json:"allocation_id,omitempty"`
}

func (o *OvertimeRecord) Prepare() error {

	if err := o.validate(); err != nil {
		return err
	}

	return nil
}

func (o *OvertimeRecord) validate() error {

	if o.StartTime.After(o.EndTime) && o.EndTime.Before(o.StartTime) {
		return errors.New("horário inválido")
	}

	return nil
}
