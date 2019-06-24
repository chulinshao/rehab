package models

import (
	"time"
)

type GoodsImg struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	GoodsCode  string    `xorm:"VARCHAR(40)"`
	OutId      string    `xorm:"comment('外部ID') VARCHAR(50)"`
	Source     string    `xorm:"comment('图片来源') VARCHAR(10)"`
	Url        string    `xorm:"comment('图片链接地址') VARCHAR(500)"`
	Thumbnail  string    `xorm:"comment('图片缩略图链接地址') VARCHAR(500)"`
	Medium     string    `xorm:"comment('中号大小图片链接地址') VARCHAR(500)"`
	Combine    string    `xorm:"comment('组合图片链接地址') VARCHAR(500)"`
	CreateTime time.Time `xorm:"comment('图片创建时间') TIMESTAMP"`
}
