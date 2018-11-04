package models

import (
	"time"
	"os"
)

type DirInfo struct {
	Id int64 `xorm:"pk autoincr notnull"`
	Path string `xorm:"varchar(100) notnull unique 'path'"`
	IsDir bool
	Length int64
	ModTime time.Time
	Owner string
}

func NewDir( fi os.FileInfo, path string ) *DirInfo{
	return &DirInfo{
		Path:path+fi.Name(),
		IsDir:fi.IsDir(),
		Length:fi.Size(),
		ModTime:fi.ModTime(),
		Owner:"nil",
	}
}