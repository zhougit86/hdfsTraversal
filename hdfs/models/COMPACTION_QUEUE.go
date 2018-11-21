package models

type CompactionQueue struct {
	CqId            int64  `xorm:"not null pk BIGINT(20)"`
	CqDatabase      string `xorm:"not null VARCHAR(128)"`
	CqTable         string `xorm:"not null VARCHAR(128)"`
	CqPartition     string `xorm:"VARCHAR(767)"`
	CqState         string `xorm:"not null CHAR(1)"`
	CqType          string `xorm:"not null CHAR(1)"`
	CqWorkerId      string `xorm:"VARCHAR(128)"`
	CqStart         int64  `xorm:"BIGINT(20)"`
	CqRunAs         string `xorm:"VARCHAR(128)"`
	CqHighestTxnId  int64  `xorm:"BIGINT(20)"`
	CqMetaInfo      []byte `xorm:"VARBINARY(2048)"`
	CqHadoopJobId   string `xorm:"VARCHAR(32)"`
	CqTblproperties string `xorm:"VARCHAR(2048)"`
}
