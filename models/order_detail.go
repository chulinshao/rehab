package models

import (
	"time"
)

type OrderDetail struct {
	Code            int64     `xorm:"not null pk autoincr comment('主键') BIGINT(20)"`
	OrderCode       string    `xorm:"not null comment('订单编码') VARCHAR(20)"`
	SourceOrderCode string    `xorm:"not null comment('来源订单编码') VARCHAR(40)"`
	GoodsCode       string    `xorm:"not null comment('商品编码') VARCHAR(40)"`
	GoodsName       string    `xorm:"comment('商品名称') VARCHAR(100)"`
	Count           string    `xorm:"not null comment('数量') DECIMAL(10)"`
	Price           string    `xorm:"not null comment('价格') DECIMAL(10)"`
	Remark          string    `xorm:"comment('备注') VARCHAR(100)"`
	CreateTime      time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('导入时间') TIMESTAMP"`
}
