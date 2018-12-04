package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"yonghui.com/hdfsTraver/hdfs/models"
	"bytes"
	"os/exec"
	"sync"
	"time"
)

var dirChan chan *xorm.Engine

func exec_shell(s *models.SyncItemOds,ngin *xorm.Engine) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("java", "-jar" ,"/home/zhouxiaogang/tbdsUncompress-1.0-SNAPSHOT.jar" , s.Path)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out


	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()


	if err!=nil{
		fmt.Println(fmt.Errorf("%s:%s",s.Path,err))
		dirChan <- ngin
		return out.String(), fmt.Errorf("%s:%s",s.Path,err)
	}
	ngin.Table(new(models.SyncItemOds)).Id(s.Id).Update(map[string]interface{}{"stage":2})
	dirChan <- ngin
	return out.String(),nil

}

func oneLoop() {
	engineRecord, error := xorm.NewEngine("mysql", "root:DataLake_Yonghui1@tcp(10.216.155.15:3306)/migration?charset=utf8")
	defer engineRecord.Close()
	if error!=nil{
		fmt.Printf("create db conn error:%s\n",error)
	}
	pEveryOne := make([]*models.SyncItemOds, 0)
	err := engineRecord.Where("mission_type = ? and stage=?",0,1).Find(&pEveryOne)
	if err!=nil{
		fmt.Println(err)
	}

	dirChan = make(chan *xorm.Engine,6)
	for i:=0;i<6;i++{
		engineRecord, error := xorm.NewEngine("mysql", "root:DataLake_Yonghui1@tcp(10.216.155.15:3306)/migration?charset=utf8")
		if error!=nil{
			panic(fmt.Sprintf("create db conn error:%s\n",error) )
		}
		dirChan <- engineRecord
	}

	for _,v :=range pEveryOne{
		fmt.Println(v)
		ngin:=<-dirChan
		go exec_shell(v,ngin)
	}

	wg :=sync.WaitGroup{}
	wg.Add(6)
	go func() {
		for {
			select {
			case eg:=<-dirChan:
				eg.Close()
				wg.Done()
			}
		}
	}()
	wg.Wait()
}

func main(){
	for{
		fmt.Printf("start loop: %s\n",time.Now())
		oneLoop()
		fmt.Printf("stop loop:  %s\n",time.Now())
		time.Sleep(20*time.Minute)
	}
}

