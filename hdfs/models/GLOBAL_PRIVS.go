package models

type GlobalPrivs struct {
	UserGrantId   int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime    int    `xorm:"not null INT(11)"`
	GrantOption   int    `xorm:"not null SMALLINT(6)"`
	Grantor       string `xorm:"unique(GLOBALPRIVILEGEINDEX) VARCHAR(128)"`
	GrantorType   string `xorm:"unique(GLOBALPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalName string `xorm:"unique(GLOBALPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalType string `xorm:"unique(GLOBALPRIVILEGEINDEX) VARCHAR(128)"`
	UserPriv      string `xorm:"unique(GLOBALPRIVILEGEINDEX) VARCHAR(128)"`
}
