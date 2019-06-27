package models

import (
	"time"
)

type DoctorAccount struct {
	Code               string    `xorm:"not null pk VARCHAR(20)" json:"code"`
	DoctorCode         string    `xorm:"not null VARCHAR(20)" json:"doctorCode"`
	Balance            string    `xorm:"not null DECIMAL(11,2)" json:"balance"`
	AlipayAccount      string    `xorm:"VARCHAR(50)" json:"alipayAccount"`
	CreateTime         time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP" json:"createTime"`
	AlipayModiftyTime  time.Time `xorm:"TIMESTAMP" json:"alipayModiftyTime"`
	BalanceModiftyTime time.Time `xorm:"TIMESTAMP" json:"balanceModiftyTime"`
}
