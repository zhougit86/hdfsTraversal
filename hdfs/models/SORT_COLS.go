package models

type SortCols struct {
	SdId       int64  `xorm:"not null pk index BIGINT(20)"`
	ColumnName string `xorm:"VARCHAR(767)"`
	Order      int    `xorm:"not null INT(11)"`
	IntegerIdx int    `xorm:"not null pk INT(11)"`
}
