package models

import (
	"time"
)

type OrderClaimLog struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderCode        string    `xorm:"not null comment('订单编码') VARCHAR(50)"`
	BeforeDoctorCode string    `xorm:"VARCHAR(40)"`
	BeforeDoctorTime time.Time `xorm:"TIMESTAMP"`
	AfterDoctorCode  string    `xorm:"not null VARCHAR(40)"`
	AfterDoctorTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
