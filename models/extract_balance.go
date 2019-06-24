package models

import (
	"time"
)

type ExtractBalance struct {
	Id              int64     `xorm:"pk autoincr BIGINT(40)"`
	DoctorCode      string    `xorm:"not null VARCHAR(20)"`
	DoctorName      string    `xorm:"not null VARCHAR(20)"`
	ExtractAmount   string    `xorm:"not null DECIMAL(11,2)"`
	CreateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Status          string    `xorm:"not null comment('申请初始状态、已提现') VARCHAR(10)"`
	HandlerUserCode string    `xorm:"VARCHAR(20)"`
	HandlerTime     time.Time `xorm:"TIMESTAMP"`
}
