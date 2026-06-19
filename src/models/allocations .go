package models

import "time"

type Allocations struct {
	Id                 uint64    `json:"id,omitempty"`
	ContractEmployeeId uint64    `json:"employee_contract_id,omitempty"`
	ContractCompanyId  uint64    `json:"company_contract_id,omitempty"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
}
