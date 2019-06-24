package models

type CommissionConfig struct {
	Id         int64  `xorm:"pk autoincr BIGINT(20)"`
	Name       string `xorm:"not null comment('名称') VARCHAR(40)"`
	Type       string `xorm:"comment('类型') VARCHAR(40)"`
	Key        string `xorm:"not null comment('键') VARCHAR(40)"`
	Value      string `xorm:"not null comment('值') VARCHAR(40)"`
	SortNum    int    `xorm:"not null INT(11)"`
	Remark     string `xorm:"comment('备注') VARCHAR(255)"`
	ModfiyCode string `xorm:"comment('最后修改人') VARCHAR(40)"`
}
