package models

import (
	"time"
)

type EvaluateLike struct {
	Id         int64     `xorm:"pk autoincr comment('id') BIGINT(20)"`
	EvaluateId int64     `xorm:"not null comment('点评ID') BIGINT(20)"`
	DoctorCode string    `xorm:"not null comment('喜欢人') VARCHAR(40)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
}
