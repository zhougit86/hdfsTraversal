package models

type Roles struct {
	RoleId     int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime int    `xorm:"not null INT(11)"`
	OwnerName  string `xorm:"VARCHAR(128)"`
	RoleName   string `xorm:"unique VARCHAR(128)"`
}
