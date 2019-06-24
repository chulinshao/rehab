package models

import (
	"time"
)

type GoodsSku struct {
	Id               int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	GoodsCode        string    `xorm:"not null comment('商品编码') VARCHAR(40)"`
	SkuJson          string    `xorm:"comment('有赞sku JSON字符串') VARCHAR(255)"`
	Name             string    `xorm:"comment('展示名称') VARCHAR(255)"`
	GoodsOutId       string    `xorm:"VARCHAR(40)"`
	SkuId            string    `xorm:"comment('有赞skuID') VARCHAR(40)"`
	SkuUniqueCode    string    `xorm:"comment('有赞sku唯一编码') VARCHAR(40)"`
	CreateTime       time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Quantity         int       `xorm:"comment('数量') INT(11)"`
	Price            string    `xorm:"comment('价格') DECIMAL(10,2)"`
	ModifyTime       time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	SoldNum          int       `xorm:"comment('Sku的销量') INT(11)"`
	WithHoldQuantity int       `xorm:"comment('商品在付款减库存的状态下，该Sku上未付款的订单数量') INT(11)"`
}
