package models

type Idxs struct {
	IndexId           int64  `xorm:"not null pk BIGINT(20)"`
	CreateTime        int    `xorm:"not null INT(11)"`
	DeferredRebuild   int    `xorm:"not null BIT(1)"`
	IndexHandlerClass string `xorm:"VARCHAR(4000)"`
	IndexName         string `xorm:"unique(UNIQUEINDEX) VARCHAR(128)"`
	IndexTblId        int64  `xorm:"index BIGINT(20)"`
	LastAccessTime    int    `xorm:"not null INT(11)"`
	OrigTblId         int64  `xorm:"index unique(UNIQUEINDEX) BIGINT(20)"`
	SdId              int64  `xorm:"index BIGINT(20)"`
}
