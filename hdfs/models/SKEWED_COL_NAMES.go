package models

type SkewedColNames struct {
	SdId          int64  `xorm:"not null pk index BIGINT(20)"`
	SkewedColName string `xorm:"VARCHAR(256)"`
	IntegerIdx    int    `xorm:"not null pk INT(11)"`
}
