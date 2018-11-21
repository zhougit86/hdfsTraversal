package models

type KeyConstraints struct {
	ChildCdId          int64  `xorm:"BIGINT(20)"`
	ChildIntegerIdx    int    `xorm:"INT(11)"`
	ChildTblId         int64  `xorm:"BIGINT(20)"`
	ParentCdId         int64  `xorm:"not null BIGINT(20)"`
	ParentIntegerIdx   int    `xorm:"not null INT(11)"`
	ParentTblId        int64  `xorm:"not null index BIGINT(20)"`
	Position           int64  `xorm:"not null pk BIGINT(20)"`
	ConstraintName     string `xorm:"not null pk VARCHAR(400)"`
	ConstraintType     int    `xorm:"not null SMALLINT(6)"`
	UpdateRule         int    `xorm:"SMALLINT(6)"`
	DeleteRule         int    `xorm:"SMALLINT(6)"`
	EnableValidateRely int    `xorm:"not null SMALLINT(6)"`
}
