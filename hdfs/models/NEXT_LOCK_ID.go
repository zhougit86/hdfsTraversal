package models

type NextLockId struct {
	NlNext int64 `xorm:"not null BIGINT(20)"`
}
