package models

type SkewedColValueLocMap struct {
	SdId            int64  `xorm:"not null pk index BIGINT(20)"`
	StringListIdKid int64  `xorm:"not null pk index BIGINT(20)"`
	Location        string `xorm:"VARCHAR(4000)"`
}
