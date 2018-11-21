package models

type Partitions struct {
	PartId         int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime     int    `xorm:"not null INT(11)"`
	LastAccessTime int    `xorm:"not null INT(11)"`
	PartName       string `xorm:"unique(UNIQUEPARTITION) VARCHAR(767)"`
	SdId           int64  `xorm:"index BIGINT(20)"`
	TblId          int64  `xorm:"index unique(UNIQUEPARTITION) BIGINT(20)"`
}
