package models

import (
	"time"
)

type Goods struct {
	Code           string    `xorm:"not null pk comment('商品编码') VARCHAR(40)"`
	Barcode        string    `xorm:"comment('条形码') VARCHAR(50)"`
	Name           string    `xorm:"not null comment('商品名称') VARCHAR(255)"`
	RetailPrice    string    `xorm:"comment('零售价') DECIMAL(11,2)"`
	BuyingPrice    string    `xorm:"comment('进货价') DECIMAL(11,2)"`
	GrossProfit    string    `xorm:"comment('毛利额') DECIMAL(11,2)"`
	GrossMargins   string    `xorm:"comment('毛利率') DECIMAL(11,2)"`
	Trademark      string    `xorm:"comment('商标') VARCHAR(50)"`
	Manufacturer   string    `xorm:"comment('生产厂家') VARCHAR(100)"`
	Supplier       string    `xorm:"comment('供应商') VARCHAR(50)"`
	SalesClassCode string    `xorm:"comment('销售分类编码') VARCHAR(20)"`
	RoyaltyRatio   string    `xorm:"comment('提成比例') DECIMAL(11,2)"`
	IsDefault      string    `xorm:"default '0' comment('提成比例是否自定义，1：是，0：否') VARCHAR(1)"`
	IntroduceTime  time.Time `xorm:"comment('引进时间') TIMESTAMP"`
	SellPoint      string    `xorm:"comment('商品卖点') VARCHAR(200)"`
	SellUrl        string    `xorm:"comment('商铺链接') VARCHAR(500)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Creator        string    `xorm:"comment('创建人') VARCHAR(40)"`
	ModifyTime     time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	Modifier       string    `xorm:"comment('修改人') VARCHAR(40)"`
	OutId          string    `xorm:"comment('外部ID') VARCHAR(40)"`
	Source         string    `xorm:"comment('来源') VARCHAR(40)"`
	Remark         string    `xorm:"comment('商品描述') VARCHAR(500)"`
	OriginPrice    string    `xorm:"comment('原价格，可自定义格式') VARCHAR(50)"`
	ShareUrl       string    `xorm:"comment('分享出去的URl') VARCHAR(500)"`
	PicUrl         string    `xorm:"comment('主图片地址') VARCHAR(500)"`
	PicThumbUrl    string    `xorm:"comment('主图片缩略图片地址') VARCHAR(500)"`
	SoldNum        int       `xorm:"comment('总销量') INT(11)"`
	IsListing      string    `xorm:"comment('上下架状态') VARCHAR(10)"`
	SortNum        int       `xorm:"comment('排序字段') INT(11)"`
	RefGoodsCode   string    `xorm:"comment('来源商品编码') VARCHAR(20)"`
}
