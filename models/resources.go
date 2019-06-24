package models

type Resources struct {
	Id   int64  `xorm:"pk autoincr BIGINT(20)"`
	Name string `xorm:"not null comment('资源名称') VARCHAR(50)"`
	Type string `xorm:"not null comment('资源类型') VARCHAR(10)"`
	Path string `xorm:"not null comment('资源路径') VARCHAR(255)"`
}
