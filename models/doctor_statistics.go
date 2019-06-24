package models

import (
	"time"
)

type DoctorStatistics struct {
	Id             int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	DoctorCode     string    `xorm:"not null comment('医生编码') unique VARCHAR(40)"`
	EvaluateCount  int       `xorm:"not null default 0 comment('点评数') INT(11)"`
	LikeCount      int       `xorm:"not null default 0 comment('喜欢数') INT(11)"`
	ReadCount      int       `xorm:"not null default 0 comment('阅读数') INT(11)"`
	AttentionCount int       `xorm:"not null default 0 comment('关注数量') INT(11)"`
	FansCount      int       `xorm:"not null default 0 comment('粉丝数量') INT(11)"`
	CountTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后统计时间') TIMESTAMP"`
}
