package models

type Serdes struct {
	SerdeId int64  `xorm:"not null pk BIGINT(20)"`
	Name    string `xorm:"VARCHAR(128)"`
	Slib    string `xorm:"VARCHAR(4000)"`
}
