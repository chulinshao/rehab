package models

type Dict struct {
	Code         string `xorm:"not null pk VARCHAR(20)"`
	Name         string `xorm:"not null VARCHAR(20)"`
	Level        int    `xorm:"comment('层级，从1开始') INT(255)"`
	ClassifyCode string `xorm:"VARCHAR(20)"`
	ParentCode   string `xorm:"VARCHAR(20)"`
	PathName     string `xorm:"comment('名称的全路径') VARCHAR(255)"`
	IsLeaf       string `xorm:"comment('是否叶子结点，0-是，1-否') VARCHAR(1)"`
	Remark       string `xorm:"VARCHAR(50)"`
}
