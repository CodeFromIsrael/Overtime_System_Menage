package models

import "time"

type ContractEmployee struct {
	Id               uint64    `json:"id,omitempty"`
	AdmissionDate    time.Time `json:"admission_date"`
	TerminationDate  time.Time `json:"terminatio_date"`
	ServiceCompanyId uint64    `json:"service_company_id"`
	UserId           uint64    `json:"user_id,omitempty"`
	Status           string    `json:"status,omitempty"`
	Position         string    `json:"position,omitempty"`
}
