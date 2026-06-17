package models

import "time"

type ContractCompany struct {
	Id               uint64    `json:"id,omitempty"`
	ServiceCompanyId uint64    `json:"service_company_id,omitempty"`
	ClientCompanyId  uint64    `json:"client_company_id,omitempty"`
	StartTime        time.Time `json:"start_date"`
	EndTime          time.Time `json:"end_date"`
	State            string    `json:"status,omitempty"`
}
