package models

import "time"

type Logger struct {
	id         uint64
	UserId     uint64
	Date       time.Time
	Request    string
	MethodHttp string
	Body       []byte
}
