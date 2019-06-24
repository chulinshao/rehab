package models

import (
	"time"
)

type DoctorAttention struct {
	Id                  int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	DoctorCode          string    `xorm:"not null comment('医生编码') VARCHAR(40)"`
	AttentionDoctorCode string    `xorm:"not null comment('关注医生') VARCHAR(40)"`
	AttentionTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('关注时间') TIMESTAMP"`
	ModifyTime          time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	IsDel               string    `xorm:"not null default 'N' comment('是否删除') VARCHAR(1)"`
}
