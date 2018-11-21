package models

type NotificationLog struct {
	NlId          int64  `xorm:"not null pk BIGINT(20)"`
	EventId       int64  `xorm:"not null BIGINT(20)"`
	EventTime     int    `xorm:"not null INT(11)"`
	EventType     string `xorm:"not null VARCHAR(32)"`
	DbName        string `xorm:"VARCHAR(128)"`
	TblName       string `xorm:"VARCHAR(256)"`
	Message       string `xorm:"LONGTEXT"`
	MessageFormat string `xorm:"VARCHAR(16)"`
}
