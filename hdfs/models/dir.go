package models

import (
	"time"
)

type Dir_Inf struct {
	ID int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Path string `gorm:"UNIQUE;NOT NULL"`
	IsDir bool
	Length int64
	ModTime time.Time
	Owner string
}


