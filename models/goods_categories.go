package models

import (
	"time"
)

type GoodsCategories struct {
	Id           int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	GoodsCode    string    `xorm:"not null comment('商品编码') VARCHAR(40)"`
	CategoriesId int64     `xorm:"not null comment('类目ID') BIGINT(20)"`
	OutId        string    `xorm:"comment('外部类目ID') VARCHAR(50)"`
	Source       string    `xorm:"VARCHAR(10)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	ModfiyTime   time.Time `xorm:"comment('修改时间') TIMESTAMP"`
}
