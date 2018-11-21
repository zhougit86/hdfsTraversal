package models

type Dbs struct {
	DbId          int64  `xorm:"not null pk BIGINT(20)"`
	Desc          string `xorm:"VARCHAR(4000)"`
	DbLocationUri string `xorm:"not null VARCHAR(4000)"`
	Name          string `xorm:"unique VARCHAR(128)"`
	OwnerName     string `xorm:"VARCHAR(128)"`
	OwnerType     string `xorm:"VARCHAR(10)"`
}
