package models

import (
	"time"
)

type Sms struct {
	Id               int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Mobile           string    `xorm:"not null comment('手机号码') VARCHAR(11)"`
	SmsCode          string    `xorm:"not null comment('短信验证码') VARCHAR(8)"`
	CreateTime       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	InvalidTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('失效时间') TIMESTAMP"`
	Source           string    `xorm:"not null comment('来源') VARCHAR(10)"`
	Status           string    `xorm:"not null comment('状态：0生效，1失效') VARCHAR(1)"`
	ReturnStatus     string    `xorm:"comment('发送状态') VARCHAR(20)"`
	ReturnStatusDesc string    `xorm:"comment('返回状态码描述') VARCHAR(20)"`
	ReturnId         string    `xorm:"comment('发送回执ID') VARCHAR(40)"`
}
