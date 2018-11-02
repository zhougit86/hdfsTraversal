package models

import (
	"time"
	"os"
)

type Dir_Inf struct {
	ID int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Path string `gorm:"UNIQUE;NOT NULL"`
	IsDir bool
	Length int64
	ModTime time.Time
	Owner string
}

func NewDir( fi os.FileInfo, path string ) *Dir_Inf{
	return &Dir_Inf{
		Path:path+fi.Name(),
		IsDir:fi.IsDir(),
		Length:fi.Size(),
		ModTime:fi.ModTime(),
		Owner:"nil",
	}
}
