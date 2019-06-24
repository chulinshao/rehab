package models

import (
	"time"
)

type HospitalTreatRoom struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	TreatRoomName    string    `xorm:"not null comment('治疗室名称') VARCHAR(50)"`
	DepId            int64     `xorm:"not null comment('关联的科室id') BIGINT(20)"`
	CreateTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateDoctorCode string    `xorm:"not null comment('创建医生编码') VARCHAR(40)"`
}
