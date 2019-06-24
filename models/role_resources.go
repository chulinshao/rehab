package models

type RoleResources struct {
	ResourcesId int64 `xorm:"not null pk BIGINT(20)"`
	RoleId      int64 `xorm:"not null pk BIGINT(20)"`
}
