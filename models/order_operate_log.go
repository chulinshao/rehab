package models

import (
	"time"
)

type OrderOperateLog struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	OrderCode   string    `xorm:"not null comment('订单编码') VARCHAR(50)"`
	DoctorCode  string    `xorm:"comment('医生') VARCHAR(40)"`
	UserCode    string    `xorm:"comment('用户') VARCHAR(40)"`
	OperateType string    `xorm:"comment('操作类型') VARCHAR(40)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
