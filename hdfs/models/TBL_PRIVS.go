package models

type TblPrivs struct {
	TblGrantId    int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime    int    `xorm:"not null INT(11)"`
	GrantOption   int    `xorm:"not null SMALLINT(6)"`
	Grantor       string `xorm:"index(TABLEPRIVILEGEINDEX) VARCHAR(128)"`
	GrantorType   string `xorm:"index(TABLEPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalName string `xorm:"index(TABLEPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalType string `xorm:"index(TABLEPRIVILEGEINDEX) VARCHAR(128)"`
	TblPriv       string `xorm:"index(TABLEPRIVILEGEINDEX) VARCHAR(128)"`
	TblId         int64  `xorm:"index(TABLEPRIVILEGEINDEX) index BIGINT(20)"`
}
