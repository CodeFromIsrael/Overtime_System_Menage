package models

type UserAndRole struct {
	User Users `json:"users"`
	Role Roles `json:"roles"`
}
