package models

type Role struct {
	Id   int64  `xorm:"pk autoincr BIGINT(20)"`
	Code string `xorm:"comment('角色编码') VARCHAR(40)"`
	Name string `xorm:"comment('角色名称') VARCHAR(40)"`
}
