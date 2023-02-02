package models

import (
	"time"
)

type QueryRecord struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	ToUserId   int       `xorm:"INT"`
	UpdateTime time.Time `xorm:"DATETIME"`
	Response   string    `xorm:"comment('?: No answer yet
yes: yes
no: no
later: later') VARCHAR(255)"`
}
