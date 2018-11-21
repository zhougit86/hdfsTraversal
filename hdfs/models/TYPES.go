package models

type Types struct {
	TypesId  int64  `xorm:"not null pk BIGINT(20)"`
	TypeName string `xorm:"unique VARCHAR(128)"`
	Type1    string `xorm:"VARCHAR(767)"`
	Type2    string `xorm:"VARCHAR(767)"`
}
