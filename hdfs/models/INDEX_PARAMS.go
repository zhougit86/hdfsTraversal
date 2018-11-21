package models

type IndexParams struct {
	IndexId    int64  `xorm:"not null pk index BIGINT(20)"`
	ParamKey   string `xorm:"not null pk VARCHAR(256)"`
	ParamValue string `xorm:"VARCHAR(4000)"`
}
