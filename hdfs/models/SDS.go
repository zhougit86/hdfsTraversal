package models

type Sds struct {
	SdId                     int64  `xorm:"not null pk BIGINT(20)"`
	CdId                     int64  `xorm:"index BIGINT(20)"`
	InputFormat              string `xorm:"VARCHAR(4000)"`
	IsCompressed             int    `xorm:"not null BIT(1)"`
	IsStoredassubdirectories int    `xorm:"not null BIT(1)"`
	Location                 string `xorm:"VARCHAR(4000)"`
	NumBuckets               int    `xorm:"not null INT(11)"`
	OutputFormat             string `xorm:"VARCHAR(4000)"`
	SerdeId                  int64  `xorm:"index BIGINT(20)"`
}

func (s *Sds) TableName() string{
	return "SDS"
}