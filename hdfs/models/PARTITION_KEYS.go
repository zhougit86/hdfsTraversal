package models

type PartitionKeys struct {
	TblId       int64  `xorm:"not null pk index BIGINT(20)"`
	PkeyComment string `xorm:"VARCHAR(4000)"`
	PkeyName    string `xorm:"not null pk VARCHAR(128)"`
	PkeyType    string `xorm:"not null VARCHAR(767)"`
	IntegerIdx  int    `xorm:"not null INT(11)"`
}
