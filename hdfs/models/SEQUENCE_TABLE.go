package models

type SequenceTable struct {
	SequenceName string `xorm:"not null pk VARCHAR(255)"`
	NextVal      int64  `xorm:"not null BIGINT(20)"`
}
