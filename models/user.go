package models

import (
	"time"
)

type User struct {
	Code          string    `xorm:"not null pk VARCHAR(20)"`
	LoginCode     string    `xorm:"not null VARCHAR(20)"`
	Name          string    `xorm:"not null VARCHAR(20)"`
	Password      string    `xorm:"not null VARCHAR(100)"`
	Mobile        string    `xorm:"VARCHAR(11)"`
	CreateTime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CreateName    string    `xorm:"not null VARCHAR(20)"`
	LastLoginTime time.Time `xorm:"TIMESTAMP"`
	LastLoginIp   string    `xorm:"VARCHAR(50)"`
	Status        string    `xorm:"VARCHAR(10)"`
}
