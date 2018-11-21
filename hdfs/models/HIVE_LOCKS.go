package models

type HiveLocks struct {
	HlLockExtId      int64  `xorm:"not null pk BIGINT(20)"`
	HlLockIntId      int64  `xorm:"not null pk BIGINT(20)"`
	HlTxnid          int64  `xorm:"index index BIGINT(20)"`
	HlDb             string `xorm:"not null VARCHAR(128)"`
	HlTable          string `xorm:"VARCHAR(128)"`
	HlPartition      string `xorm:"VARCHAR(767)"`
	HlLockState      string `xorm:"not null CHAR(1)"`
	HlLockType       string `xorm:"not null CHAR(1)"`
	HlLastHeartbeat  int64  `xorm:"not null BIGINT(20)"`
	HlAcquiredAt     int64  `xorm:"BIGINT(20)"`
	HlUser           string `xorm:"not null VARCHAR(128)"`
	HlHost           string `xorm:"not null VARCHAR(128)"`
	HlHeartbeatCount int    `xorm:"INT(11)"`
	HlAgentInfo      string `xorm:"VARCHAR(128)"`
	HlBlockedbyExtId int64  `xorm:"BIGINT(20)"`
	HlBlockedbyIntId int64  `xorm:"BIGINT(20)"`
}
