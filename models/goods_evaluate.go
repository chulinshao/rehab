package models

import (
	"time"
)

type GoodsEvaluate struct {
	Id           int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	GoodsCode    string    `xorm:"not null comment('商品编码') VARCHAR(40)"`
	PatientUsed  string    `xorm:"comment('是否为患者使用过产品') VARCHAR(1)"`
	ApplyDisease string    `xorm:"comment('适用的疾病') VARCHAR(255)"`
	ApplySymptom string    `xorm:"comment('适用症状') VARCHAR(255)"`
	ApplyStage   string    `xorm:"comment('适用阶段') VARCHAR(255)"`
	ApplyPerson  string    `xorm:"comment('适用人员') VARCHAR(255)"`
	GoodsType    string    `xorm:"comment('型号/颜色分类') VARCHAR(2000)"`
	Remark       string    `xorm:"comment('描述') VARCHAR(2000)"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('点评时间') TIMESTAMP"`
	DoctorCode   string    `xorm:"not null comment('点评医生') VARCHAR(40)"`
	ReadCount    int       `xorm:"not null default 0 comment('阅读数量') INT(11)"`
	ReplyCount   int       `xorm:"not null default 0 comment('回复数量') INT(11)"`
	LikeCount    int       `xorm:"not null default 0 comment('喜欢数量') INT(11)"`
	ShareCount   int       `xorm:"not null default 0 comment('分享数') INT(11)"`
	IsDel        string    `xorm:"not null default 'N' comment('是否删除') VARCHAR(1)"`
}
