package models

import (
	"time"
)

type DoctorAccount struct {
	Code               string    `xorm:"not null pk VARCHAR(20)"`
	DoctorCode         string    `xorm:"not null VARCHAR(20)"`
	Balance            string    `xorm:"not null DECIMAL(11,2)"`
	AlipayAccount      string    `xorm:"VARCHAR(50)"`
	CreateTime         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	AlipayModiftyTime  time.Time `xorm:"TIMESTAMP"`
	BalanceModiftyTime time.Time `xorm:"TIMESTAMP"`
}
