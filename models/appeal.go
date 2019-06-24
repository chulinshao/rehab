package models

import (
	"time"
)

type Appeal struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	DoctorCode     string    `xorm:"not null VARCHAR(20)"`
	DoctorName     string    `xorm:"not null VARCHAR(20)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Remark         string    `xorm:"VARCHAR(50)"`
	OrderCode      string    `xorm:"not null VARCHAR(20)"`
	Status         string    `xorm:"not null comment('待处理，已处理') VARCHAR(10)"`
	HandleUserCode string    `xorm:"VARCHAR(20)"`
	Result         string    `xorm:"comment('通过、不通过') VARCHAR(6)"`
	ResultRemark   string    `xorm:"VARCHAR(100)"`
	HandleTime     time.Time `xorm:"TIMESTAMP"`
}
