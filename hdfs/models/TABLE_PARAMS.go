package models

type TableParams struct {
	TblId      int64  `xorm:"not null pk index BIGINT(20)"`
	ParamKey   string `xorm:"not null pk VARCHAR(256)"`
	ParamValue string `xorm:"VARCHAR(4000)"`
}
