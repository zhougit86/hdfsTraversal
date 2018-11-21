package models

type Funcs struct {
	FuncId     int64  `xorm:"not null pk BIGINT(20)"`
	ClassName  string `xorm:"VARCHAR(4000)"`
	CreateTime int    `xorm:"not null INT(11)"`
	DbId       int64  `xorm:"index unique(UNIQUEFUNCTION) BIGINT(20)"`
	FuncName   string `xorm:"unique(UNIQUEFUNCTION) VARCHAR(128)"`
	FuncType   int    `xorm:"not null INT(11)"`
	OwnerName  string `xorm:"VARCHAR(128)"`
	OwnerType  string `xorm:"VARCHAR(10)"`
}
