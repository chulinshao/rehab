package models

import (
	"time"
)

type DoctorAccountLog struct {
	Id         int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	DoctorCode string    `xorm:"not null comment('医生编码') VARCHAR(40)"`
	OrderCode  string    `xorm:"VARCHAR(50)"`
	Balance    string    `xorm:"not null comment('变动金额') DECIMAL(11,2)"`
	LogType    string    `xorm:"not null default 'ORDER' comment('日志类型') VARCHAR(10)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Remark     string    `xorm:"VARCHAR(100)"`
}
