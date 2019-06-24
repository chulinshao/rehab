package models

import (
	"time"
)

type DoctorOrder struct {
	DoctorCode     string    `xorm:"not null pk comment('医生编码') VARCHAR(20)"`
	OrderCode      string    `xorm:"not null pk comment('订单编码') VARCHAR(20)"`
	IsAppeal       string    `xorm:"not null comment('是否申诉') VARCHAR(1)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	LastAppealTime time.Time `xorm:"comment('最后申请时间') TIMESTAMP"`
}
