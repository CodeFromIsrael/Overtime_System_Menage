package responses

import "overtime_system_menagement/src/models"

type UserAndRole struct {
	User models.Users `json:"users"`
	Role models.Roles `json:"roles"`
}
