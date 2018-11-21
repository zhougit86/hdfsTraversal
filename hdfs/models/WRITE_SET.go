package models

type WriteSet struct {
	WsDatabase      string `xorm:"not null VARCHAR(128)"`
	WsTable         string `xorm:"not null VARCHAR(128)"`
	WsPartition     string `xorm:"VARCHAR(767)"`
	WsTxnid         int64  `xorm:"not null BIGINT(20)"`
	WsCommitId      int64  `xorm:"not null BIGINT(20)"`
	WsOperationType string `xorm:"not null CHAR(1)"`
}
