package models

type OvertimeTypes struct {
	Id         uint64  `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}
