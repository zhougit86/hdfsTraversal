package models

type TypeFields struct {
	TypeName   int64  `xorm:"not null pk index BIGINT(20)"`
	Comment    string `xorm:"VARCHAR(256)"`
	FieldName  string `xorm:"not null pk VARCHAR(128)"`
	FieldType  string `xorm:"not null VARCHAR(767)"`
	IntegerIdx int    `xorm:"not null INT(11)"`
}
