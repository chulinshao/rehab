package models

import (
	"time"
)

type Categories struct {
	Id         int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	OutId      string    `xorm:"comment('外部ID') VARCHAR(50)"`
	Name       string    `xorm:"comment('类目名称') VARCHAR(50)"`
	Alias      string    `xorm:"comment('别名') VARCHAR(100)"`
	TagUrl     string    `xorm:"comment('分组链接') VARCHAR(500)"`
	ShareUrl   string    `xorm:"comment('分组链接') VARCHAR(500)"`
	ItemNum    int       `xorm:"comment('商品数量') INT(11)"`
	Source     string    `xorm:"comment('分享链接') VARCHAR(10)"`
	Remark     string    `xorm:"comment('描述') VARCHAR(255)"`
	CreateTime time.Time `xorm:"TIMESTAMP"`
	ModifyTime time.Time `xorm:"TIMESTAMP"`
}
