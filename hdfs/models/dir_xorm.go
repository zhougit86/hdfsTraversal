package models

import (
	"time"
	"os"
	"github.com/colinmarc/hdfs/v2"
)

type DirInfo struct {
	Id int64 `xorm:"pk autoincr notnull"`
	Path string `xorm:"notnull unique 'path'"`
	IsDir bool
	Length int64
	ModTime time.Time
	Owner string
}

func NewDir( fi os.FileInfo, path string ) *DirInfo{
	dir:= &DirInfo{
		Path:path+fi.Name(),
		IsDir:fi.IsDir(),
		Length:fi.Size(),
		ModTime:fi.ModTime(),
	}
	if fiHDFS,ok:=fi.(*hdfs.FileInfo);ok{
		dir.Owner = fiHDFS.Owner()
	}
	return dir
}