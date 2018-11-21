package models

type FuncRu struct {
	FuncId       int64  `xorm:"not null pk BIGINT(20)"`
	ResourceType int    `xorm:"not null INT(11)"`
	ResourceUri  string `xorm:"VARCHAR(4000)"`
	IntegerIdx   int    `xorm:"not null pk INT(11)"`
}
