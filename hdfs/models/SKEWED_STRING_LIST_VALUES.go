package models

type SkewedStringListValues struct {
	StringListId    int64  `xorm:"not null pk index BIGINT(20)"`
	StringListValue string `xorm:"VARCHAR(256)"`
	IntegerIdx      int    `xorm:"not null pk INT(11)"`
}
