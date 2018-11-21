package models

type NextCompactionQueueId struct {
	NcqNext int64 `xorm:"not null BIGINT(20)"`
}
