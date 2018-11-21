package models

type CompletedTxnComponents struct {
	CtcTxnid     int64  `xorm:"BIGINT(20)"`
	CtcDatabase  string `xorm:"not null VARCHAR(128)"`
	CtcTable     string `xorm:"VARCHAR(256)"`
	CtcPartition string `xorm:"VARCHAR(767)"`
}
