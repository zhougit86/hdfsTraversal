package models

type ColumnsV2 struct {
	CdId       int64  `xorm:"not null pk index BIGINT(20)"`
	Comment    string `xorm:"VARCHAR(256)"`
	ColumnName string `xorm:"not null pk VARCHAR(767)"`
	TypeName   string `xorm:"MEDIUMTEXT"`
	IntegerIdx int    `xorm:"not null INT(11)"`
}
