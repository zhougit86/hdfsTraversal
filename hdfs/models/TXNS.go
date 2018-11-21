package models

type Txns struct {
	TxnId             int64  `xorm:"not null pk BIGINT(20)"`
	TxnState          string `xorm:"not null CHAR(1)"`
	TxnStarted        int64  `xorm:"not null BIGINT(20)"`
	TxnLastHeartbeat  int64  `xorm:"not null BIGINT(20)"`
	TxnUser           string `xorm:"not null VARCHAR(128)"`
	TxnHost           string `xorm:"not null VARCHAR(128)"`
	TxnAgentInfo      string `xorm:"VARCHAR(128)"`
	TxnHeartbeatCount int    `xorm:"INT(11)"`
	TxnMetaInfo       string `xorm:"VARCHAR(128)"`
}
