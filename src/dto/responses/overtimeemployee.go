package responses

import "overtime_system_menagement/src/models"

type OvertimeEmployee struct {
	Overtime              models.OvertimeRecord `json:"overtime_records"`
	TypeStartTimeReturned string                `json:"start_time"`
	TypeEndtimeReturned   string                `json:"end_time"`
	Employee              models.Users          `json:"users"`
}
