package models

type NotificationSequence struct {
	NniId       int64 `xorm:"not null pk BIGINT(20)"`
	NextEventId int64 `xorm:"not null BIGINT(20)"`
}
