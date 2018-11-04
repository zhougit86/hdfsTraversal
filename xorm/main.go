package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"fmt"
	"yonghui.com/hdfsTraver/hdfs/models"
	"strconv"
	"time"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func main() {

	engine, _ = xorm.NewEngine("mysql", "root:123@tcp(192.168.13.128:3306)/migration?charset=utf8")

	engine.CreateTables(models.DirInfo{})
	engine.CreateUniques(models.DirInfo{})

	ddd:=&models.DirInfo{
		Path:"ttt5",
		ModTime:time.Now(),
	}
	engine.Insert(ddd)

	ddd6:=&models.DirInfo{
		Path:"ttt6",
		ModTime:time.Now(),
	}
	engine.Insert(ddd6)

	dirList := make([]*models.DirInfo,150)
	for i:=0;i<150;i++{
		dirList[i] = &models.DirInfo{
			Path:"ttt"+strconv.Itoa(i),
		}
	}
	num,err:=engine.Insert(&dirList)
	mysqlErr:=err.(*mysql.MySQLError)
	fmt.Println(num)
	fmt.Println(mysqlErr.Number)

}