package models

import (
	"time"
)

type GoodsFill struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	OutId      string    `xorm:"VARCHAR(255)"`
	GoodsCode  string    `xorm:"not null VARCHAR(50)"`
	Title      string    `xorm:"VARCHAR(255)"`
	CreateTime time.Time `xorm:"TIMESTAMP"`
}
