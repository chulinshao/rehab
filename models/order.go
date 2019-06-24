package models

import (
	"time"
)

type Order struct {
	Code            string    `xorm:"not null pk comment('订单编码') VARCHAR(20)"`
	Source          string    `xorm:"not null comment('订单来源') VARCHAR(6)"`
	SourceOrderCode string    `xorm:"not null comment('来源订单编码') VARCHAR(50)"`
	ReceiverName    string    `xorm:"not null comment('收货人姓名') VARCHAR(20)"`
	ReceiverMobile  string    `xorm:"not null comment('收货人电话') VARCHAR(20)"`
	ReceiverAddress string    `xorm:"not null comment('收货人地址') VARCHAR(255)"`
	AmountPrice     string    `xorm:"comment('应付金额') DECIMAL(11,2)"`
	ActuallyPrice   string    `xorm:"not null comment('实付金额') DECIMAL(11,2)"`
	FreightPrice    string    `xorm:"comment('运费金额') DECIMAL(11,2)"`
	Remark          string    `xorm:"comment('订单备注') VARCHAR(100)"`
	SourceStatus    string    `xorm:"comment('订单来源状态') VARCHAR(50)"`
	Status          string    `xorm:"not null comment('订单状态：待认领、已认领') VARCHAR(40)"`
	OrderCreateTime time.Time `xorm:"comment('订单创建时间') TIMESTAMP"`
	Creator         string    `xorm:"comment('导入者编码') VARCHAR(40)"`
	CreateTime      time.Time `xorm:"comment('导入时间') TIMESTAMP"`
	DoctorCode      string    `xorm:"comment('领取医生编码') VARCHAR(20)"`
	ClaimTime       time.Time `xorm:"comment('认领时间') TIMESTAMP"`
	IsAppeal        string    `xorm:"comment('是否申诉') VARCHAR(1)"`
	LastAppealTime  time.Time `xorm:"comment('最后申诉时间') TIMESTAMP"`
	RatioMoney      string    `xorm:"comment('提成金额') DECIMAL(11,2)"`
	PublicEndTime   time.Time `xorm:"comment('公示截止时间') TIMESTAMP"`
}
