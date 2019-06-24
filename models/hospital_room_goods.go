package models

type HospitalRoomGoods struct {
	Id          int64  `xorm:"pk autoincr BIGINT(20)"`
	GoodsCode   string `xorm:"not null comment('商品编码') VARCHAR(40)"`
	TreatRoomId int64  `xorm:"not null comment('科室名称') BIGINT(20)"`
}
