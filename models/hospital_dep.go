package models

import (
	"time"
)

type HospitalDep struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	DepName          string    `xorm:"not null comment('科室名称') VARCHAR(50)"`
	HospitalCode     string    `xorm:"not null VARCHAR(40)"`
	CreateTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateDoctorCode string    `xorm:"not null comment('创建医生编码') VARCHAR(40)"`
}
