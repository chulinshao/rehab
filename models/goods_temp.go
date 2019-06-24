package models

type GoodsTemp struct {
	Code           string `xorm:"not null pk VARCHAR(50)"`
	Name           string `xorm:"VARCHAR(255)"`
	Trademark      string `xorm:"VARCHAR(255)"`
	Manufacturer   string `xorm:"VARCHAR(255)"`
	Supplier       string `xorm:"VARCHAR(255)"`
	RetailPrice    string `xorm:"DECIMAL(10,2)"`
	SalesClassCode string `xorm:"VARCHAR(10)"`
	RoyaltyRatio   string `xorm:"DECIMAL(10,2)"`
	SellPoint      string `xorm:"VARCHAR(255)"`
	SellUrl        string `xorm:"VARCHAR(255)"`
	Professional   string `xorm:"VARCHAR(2000)"`
}
