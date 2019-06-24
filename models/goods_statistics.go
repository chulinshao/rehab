package models

import (
	"time"
)

type GoodsStatistics struct {
	Id            int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	GoodsCode     string    `xorm:"not null comment('医生编码') unique VARCHAR(40)"`
	EvaluateCount int       `xorm:"not null default 0 comment('点评数') INT(11)"`
	LikeCount     int       `xorm:"not null default 0 comment('喜欢数') INT(11)"`
	ReadCount     int       `xorm:"not null default 0 comment('阅读数') INT(11)"`
	CountTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后统计时间') TIMESTAMP"`
}
