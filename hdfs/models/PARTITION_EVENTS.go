package models

type PartitionEvents struct {
	PartNameId    int64  `xorm:"not null pk BIGINT(20)"`
	DbName        string `xorm:"VARCHAR(128)"`
	EventTime     int64  `xorm:"not null BIGINT(20)"`
	EventType     int    `xorm:"not null INT(11)"`
	PartitionName string `xorm:"index VARCHAR(767)"`
	TblName       string `xorm:"VARCHAR(256)"`
}
