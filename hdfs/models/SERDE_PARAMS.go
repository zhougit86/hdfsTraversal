package models

type SerdeParams struct {
	SerdeId    int64  `xorm:"not null pk index BIGINT(20)"`
	ParamKey   string `xorm:"not null pk VARCHAR(256)"`
	ParamValue string `xorm:"MEDIUMTEXT"`
}
