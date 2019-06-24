package models

import (
	"time"
)

type SyncLog struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	TaskId        string    `xorm:"not null comment('任务ID') VARCHAR(40)"`
	OperateTyep   string    `xorm:"not null comment('操作类型') VARCHAR(10)"`
	OperateResult string    `xorm:"not null comment('操作结果') VARCHAR(10)"`
	Remark        string    `xorm:"comment('结果备注') TEXT"`
	CreateTime    time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
