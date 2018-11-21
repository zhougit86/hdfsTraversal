package models

type NextTxnId struct {
	NtxnNext int64 `xorm:"not null BIGINT(20)"`
}
