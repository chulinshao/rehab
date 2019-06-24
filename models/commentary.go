package models

import (
	"time"
)

type Commentary struct {
	Id          int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	EvaluateId  int64     `xorm:"not null comment('点评ID') index BIGINT(20)"`
	Content     string    `xorm:"not null comment('评论内容') VARCHAR(2000)"`
	ParentId    int       `xorm:"comment('上级评论内容') INT(11)"`
	DoctorCode  string    `xorm:"comment('评论医生') VARCHAR(40)"`
	ReplyCount  int       `xorm:"not null default 0 comment('回复数') INT(11)"`
	PraiseCount int       `xorm:"not null default 0 INT(11)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('评论时间') TIMESTAMP"`
	IsDel       string    `xorm:"not null default 'N' comment('是否删除') VARCHAR(1)"`
}
