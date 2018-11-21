package models

type MasterKeys struct {
	KeyId     int    `xorm:"not null pk autoincr INT(11)"`
	MasterKey string `xorm:"VARCHAR(767)"`
}
