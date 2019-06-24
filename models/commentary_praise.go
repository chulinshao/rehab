package models

import (
	"time"
)

type CommentaryPraise struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	CommentaryId     int64     `xorm:"not null index(idx_id_goods) BIGINT(20)"`
	PraiseDoctorCode string    `xorm:"not null index(idx_id_goods) VARCHAR(40)"`
	CreateTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	ModfiyTime       time.Time `xorm:"TIMESTAMP"`
	IsDel            string    `xorm:"not null default 'N' VARCHAR(1)"`
}
