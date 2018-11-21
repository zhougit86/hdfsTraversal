package models

type AuxTable struct {
	MtKey1    string `xorm:"not null pk VARCHAR(128)"`
	MtKey2    int64  `xorm:"not null pk BIGINT(20)"`
	MtComment string `xorm:"VARCHAR(255)"`
}
