package models

type PartPrivs struct {
	PartGrantId   int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime    int    `xorm:"not null INT(11)"`
	GrantOption   int    `xorm:"not null SMALLINT(6)"`
	Grantor       string `xorm:"index(PARTPRIVILEGEINDEX) VARCHAR(128)"`
	GrantorType   string `xorm:"index(PARTPRIVILEGEINDEX) VARCHAR(128)"`
	PartId        int64  `xorm:"index(PARTPRIVILEGEINDEX) index BIGINT(20)"`
	PrincipalName string `xorm:"index(PARTPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalType string `xorm:"index(PARTPRIVILEGEINDEX) VARCHAR(128)"`
	PartPriv      string `xorm:"index(PARTPRIVILEGEINDEX) VARCHAR(128)"`
}
