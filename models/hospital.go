package models

type Hospital struct {
	Code       string `xorm:"not null pk VARCHAR(20)"`
	Name       string `xorm:"not null index VARCHAR(50)"`
	Province   string `xorm:"not null VARCHAR(6)"`
	Address    string `xorm:"index VARCHAR(255)"`
	Grade      string `xorm:"not null VARCHAR(20)"`
	SpecialDep string `xorm:"index VARCHAR(255)"`
	PostCodes  string `xorm:"comment('邮编') VARCHAR(8)"`
	BedNumber  int    `xorm:"comment('床位数') INT(11)"`
	Dean       string `xorm:"comment('院长') VARCHAR(20)"`
	Phone      string `xorm:"VARCHAR(50)"`
}
