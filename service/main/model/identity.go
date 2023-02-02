package models

type Identity struct {
	Id     int    `xorm:"not null pk autoincr INT"`
	CpwxId string `xorm:"VARCHAR(255)"`
	Title  string `xorm:"VARCHAR(255)"`
}
