package models

type PartColPrivs struct {
	PartColumnGrantId int64  `xorm:"not null pk BIGINT(20)"`
	ColumnName        string `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) VARCHAR(767)"`
	CreateTime        int    `xorm:"not null INT(11)"`
	GrantOption       int    `xorm:"not null SMALLINT(6)"`
	Grantor           string `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	GrantorType       string `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	PartId            int64  `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) index BIGINT(20)"`
	PrincipalName     string `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalType     string `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	PartColPriv       string `xorm:"index(PARTITIONCOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
}
