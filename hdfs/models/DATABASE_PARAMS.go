package models

type DatabaseParams struct {
	DbId       int64  `xorm:"not null pk index BIGINT(20)"`
	ParamKey   string `xorm:"not null pk VARCHAR(180)"`
	ParamValue string `xorm:"VARCHAR(4000)"`
}
