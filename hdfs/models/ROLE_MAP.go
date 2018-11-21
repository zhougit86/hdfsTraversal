package models

type RoleMap struct {
	RoleGrantId   int64  `xorm:"not null pk BIGINT(20)"`
	AddTime       int    `xorm:"not null INT(11)"`
	GrantOption   int    `xorm:"not null SMALLINT(6)"`
	Grantor       string `xorm:"unique(USERROLEMAPINDEX) VARCHAR(128)"`
	GrantorType   string `xorm:"unique(USERROLEMAPINDEX) VARCHAR(128)"`
	PrincipalName string `xorm:"unique(USERROLEMAPINDEX) VARCHAR(128)"`
	PrincipalType string `xorm:"VARCHAR(128)"`
	RoleId        int64  `xorm:"index unique(USERROLEMAPINDEX) BIGINT(20)"`
}
