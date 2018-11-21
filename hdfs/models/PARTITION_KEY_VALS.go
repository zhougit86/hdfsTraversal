package models

type PartitionKeyVals struct {
	PartId     int64  `xorm:"not null pk index BIGINT(20)"`
	PartKeyVal string `xorm:"VARCHAR(256)"`
	IntegerIdx int    `xorm:"not null pk INT(11)"`
}
