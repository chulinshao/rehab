package models

type GoodsExtend struct {
	Code                  int64  `xorm:"not null autoincr comment('code') index BIGINT(20)"`
	GoodsCode             string `xorm:"comment('商品编码') VARCHAR(40)"`
	ProfessionalClassCode string `xorm:"comment('职业分类编码') VARCHAR(20)"`
	ProfessionalClassName string `xorm:"comment('职业分类名称') VARCHAR(50)"`
}
