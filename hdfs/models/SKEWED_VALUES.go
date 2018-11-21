package models

type SkewedValues struct {
	SdIdOid         int64 `xorm:"not null pk index BIGINT(20)"`
	StringListIdEid int64 `xorm:"not null index BIGINT(20)"`
	IntegerIdx      int   `xorm:"not null pk INT(11)"`
}
