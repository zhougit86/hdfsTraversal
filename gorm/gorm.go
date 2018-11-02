package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"time"
	//"fmt"
	"fmt"
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



func main()  {
	//db, err := gorm.Open("mysql", "root:123@192.168.13.128:3306/migration?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:123@tcp(192.168.13.128:3306)/migration?charset=utf8")
	if err != nil {
		panic(err)
	}
	db.CreateTable(&Dir_Inf{})

	nn:=Dir_Inf{ Path: "Jinzhu1", IsDir: true, Length: 48,ModTime:time.Now(),Owner:"haha"}
	fmt.Println(db.NewRecord(nn))
	er:=db.Create(&nn)
	fmt.Println(db.NewRecord(nn))
	if er.Error!=nil{
		panic("primary key is existing")
	}


	//user := Dir_Inf{ID:1, Path: "Jinzhu", IsDir: true, Length: 48,ModTime:time.Now(),Owner:"haha"}





	//fmt.Println(db.NewRecord(user))
	//
	//db.Commit()
	defer db.Close()

	//db.CreateTable(&Dir_Inf{})
	//now :=time.Now()

}
