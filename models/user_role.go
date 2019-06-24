package models

type UserRole struct {
	Id       int64  `xorm:"pk autoincr BIGINT(20)"`
	UserCode string `xorm:"VARCHAR(40)"`
	RoleId   int    `xorm:"INT(11)"`
}
