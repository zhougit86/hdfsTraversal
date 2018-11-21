package models

type TblColPrivs struct {
	TblColumnGrantId int64  `xorm:"not null pk BIGINT(20)"`
	ColumnName       string `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) VARCHAR(767)"`
	CreateTime       int    `xorm:"not null INT(11)"`
	GrantOption      int    `xorm:"not null SMALLINT(6)"`
	Grantor          string `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	GrantorType      string `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalName    string `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	PrincipalType    string `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	TblColPriv       string `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) VARCHAR(128)"`
	TblId            int64  `xorm:"index(TABLECOLUMNPRIVILEGEINDEX) index BIGINT(20)"`
}
