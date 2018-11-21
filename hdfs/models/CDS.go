package models

type Cds struct {
	CdId int64 `xorm:"not null pk BIGINT(20)"`
}
