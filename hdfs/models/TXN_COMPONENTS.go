package models

type TxnComponents struct {
	TcTxnid         int64  `xorm:"index BIGINT(20)"`
	TcDatabase      string `xorm:"not null VARCHAR(128)"`
	TcTable         string `xorm:"VARCHAR(128)"`
	TcPartition     string `xorm:"VARCHAR(767)"`
	TcOperationType string `xorm:"CHAR(1)"`
}
