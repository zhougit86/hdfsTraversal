package models

type CompletedCompactions struct {
	CcId            int64  `xorm:"not null pk BIGINT(20)"`
	CcDatabase      string `xorm:"not null VARCHAR(128)"`
	CcTable         string `xorm:"not null VARCHAR(128)"`
	CcPartition     string `xorm:"VARCHAR(767)"`
	CcState         string `xorm:"not null CHAR(1)"`
	CcType          string `xorm:"not null CHAR(1)"`
	CcWorkerId      string `xorm:"VARCHAR(128)"`
	CcStart         int64  `xorm:"BIGINT(20)"`
	CcEnd           int64  `xorm:"BIGINT(20)"`
	CcRunAs         string `xorm:"VARCHAR(128)"`
	CcHighestTxnId  int64  `xorm:"BIGINT(20)"`
	CcMetaInfo      []byte `xorm:"VARBINARY(2048)"`
	CcHadoopJobId   string `xorm:"VARCHAR(32)"`
	CcTblproperties string `xorm:"VARCHAR(2048)"`
}
