package models

import (
	"github.com/chulinshao/rehab/common"
	"github.com/shopspring/decimal"
)

type DoctorAccount struct {
	Code               string           `xorm:"not null pk VARCHAR(20)" json:"code"`
	DoctorCode         string           `xorm:"not null VARCHAR(20)" json:"doctorCode"`
	Balance            string           `xorm:"not null DECIMAL(11,2)" json:"balance"`
	AlipayAccount      string           `xorm:"VARCHAR(50)" json:"alipayAccount"`
	CreateTime         common.TimeStamp `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP" json:"createTime"`
	AlipayModiftyTime  common.TimeStamp `xorm:"TIMESTAMP" json:"alipayModiftyTime"`
	BalanceModiftyTime common.TimeStamp `xorm:"TIMESTAMP" json:"balanceModiftyTime"`

	// 是否可以提现
	Apply bool `xorm:"-" json:"apply"`

	// 点评提成
	EvaluateCommentary decimal.Decimal `xorm:"-" json:"evaluateCommentary"`

	// 推荐提成
	RefereeCommentary decimal.Decimal `xorm:"-" json:"refereeCommentary"`
}
