package models

import "time"

type OvertimeCloused struct {
	Id                uint64
	ContractCompanyId uint64
	PeriodCloused     time.Time
	State             string
	ClousedBy         uint64
	ClousedAt         time.Time
}
