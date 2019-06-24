package models

type SaleClassify struct {
	Code         string  `xorm:"not null pk comment('code') VARCHAR(20)"`
	Name         string  `xorm:"not null comment('销售分类名') VARCHAR(20)"`
	RoyaltyRatio float64 `xorm:"not null comment('提成比例') DOUBLE"`
	Remark       string  `xorm:"comment('备注') VARCHAR(100)"`
}
