package models

type Tbls struct {
	TblId            int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime       int    `xorm:"not null INT(11)"`
	DbId             int64  `xorm:"index unique(UNIQUETABLE) BIGINT(20)"`
	LastAccessTime   int    `xorm:"not null INT(11)"`
	Owner            string `xorm:"VARCHAR(767)"`
	Retention        int    `xorm:"not null INT(11)"`
	SdId             int64  `xorm:"index BIGINT(20)"`
	TblName          string `xorm:"unique(UNIQUETABLE) VARCHAR(256)"`
	TblType          string `xorm:"VARCHAR(128)"`
	ViewExpandedText string `xorm:"MEDIUMTEXT"`
	ViewOriginalText string `xorm:"MEDIUMTEXT"`
	IsRewriteEnabled int    `xorm:"not null BIT(1)"`
}

func (t *Tbls) TableName() string{
	return "TBLS"
}
