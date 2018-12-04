package models

import (
	"time"
	"fmt"
)

type SyncItemOds struct {
	Id int64 `xorm:"pk autoincr notnull"`
	//Path string `xorm:"notnull unique(nameAndTime) 'path'"`
	Path string `xorm:"notnull unique(nameAndTime) 'path'"`
	//StartTime time.Time `xorm:"notnull unique(nameAndTime) "`
	StartTime time.Time `xorm:"notnull unique(nameAndTime) "`
	Stage int
	MissionType int
}

func (s *SyncItemOds) String() string{
	return fmt.Sprintf("%s:%s:%d",s.Path,s.StartTime,s.MissionType)
}