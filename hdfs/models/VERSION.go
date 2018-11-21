package models

type Version struct {
	VerId          int64  `xorm:"not null pk BIGINT(20)"`
	SchemaVersion  string `xorm:"not null VARCHAR(127)"`
	VersionComment string `xorm:"VARCHAR(255)"`
}
