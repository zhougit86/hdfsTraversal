package models

type DbPrivs struct {
	DbGrantId     int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime    int    `xorm:"not null INT(11)"`
	DbId          int64  `xorm:"unique(DBPRIVILEGEINDEX) index BIGINT(20)"`
	GrantOption   int    `xorm:"not null SMALLINT(6)"`
	Grantor       string `xorm:"unique(DBPRIVILEGEINDEX) VARCHAR(128)"`
	GrantorType   string `xorm:"unique(DBPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalName string `xorm:"unique(DBPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalType string `xorm:"unique(DBPRIVILEGEINDEX) VARCHAR(128)"`
	DbPriv        string `xorm:"unique(DBPRIVILEGEINDEX) VARCHAR(128)"`
}
