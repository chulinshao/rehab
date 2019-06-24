package models

import (
	"time"
)

type Doctor struct {
	Code             string    `xorm:"not null pk comment('医生编码') VARCHAR(20)"`
	Name             string    `xorm:"comment('医生姓名') VARCHAR(20)"`
	Password         string    `xorm:"VARCHAR(50)"`
	Mobile           string    `xorm:"not null comment('联系电话') VARCHAR(11)"`
	IsReceiveSms     string    `xorm:"default 'Y' comment('是否愿意接收短信') VARCHAR(2)"`
	Province         string    `xorm:"comment('省份') VARCHAR(6)"`
	Openid           string    `xorm:"VARCHAR(50)"`
	HospitalCode     string    `xorm:"comment('医院编码') VARCHAR(20)"`
	HospitalName     string    `xorm:"comment('医院名称') VARCHAR(50)"`
	DepCode          string    `xorm:"comment('科室编码') VARCHAR(20)"`
	DepName          string    `xorm:"comment('科室名称') VARCHAR(20)"`
	ProfessionCode   string    `xorm:"comment('职业编码') VARCHAR(20)"`
	ProfessionName   string    `xorm:"comment('职业名称') VARCHAR(20)"`
	RecoveredCode    string    `xorm:"comment('康复方向') VARCHAR(20)"`
	RecoveredName    string    `xorm:"comment('康复方向名称') VARCHAR(20)"`
	RecoveredSubCode string    `xorm:"comment('康复子方向编码') VARCHAR(20)"`
	RecoveredSubName string    `xorm:"comment('康复子方向名称') VARCHAR(20)"`
	PortraitImg      string    `xorm:"comment('头像') VARCHAR(255)"`
	RefereeCode      string    `xorm:"comment('推荐人编码') VARCHAR(20)"`
	AuditStatus      string    `xorm:"comment('审核状态') VARCHAR(10)"`
	Status           string    `xorm:"not null default 'ENABLE' comment('医生状态') VARCHAR(10)"`
	Auditor          string    `xorm:"comment('审核人') VARCHAR(20)"`
	AuditTime        time.Time `xorm:"comment('审核时间') TIMESTAMP"`
	CreateTime       time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	ModifyTime       time.Time `xorm:"TIMESTAMP"`
}
