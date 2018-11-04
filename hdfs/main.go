package main

import(
	"github.com/colinmarc/hdfs/v2"

	"flag"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yonghui.com/hdfsTraver/hdfs/models"
	//"fmt"
	"fmt"
	"github.com/go-xorm/xorm"
	"bytes"
	"math/rand"
)

var clients chan *hdfs.Client
var objChannel chan *models.DirInfo
var endChannel chan int

var hdfsDest *string
var rootDir *string
var mysqlAddr *string
var clientNum *int
var sqlConnNum *int


func init()  {
	mysqlAddr= flag.String("sql", "root:123@tcp(192.168.13.128:3306)/migration?charset=utf8", "sql:root:123@tcp(192.168.13.128:3306)/migration?charset=utf8")
	hdfsDest= flag.String("hdfs", "192.168.13.128:9000", "The hdfs you want to connect")
	rootDir = flag.String("root" ,"/","the directory to start with")
	clientNum = flag.Int("client",3,"the number of client to hdfs")
	sqlConnNum = flag.Int("sqlconn",5,"the number of sql conn")
}

func batchPersistence(){
	var db *xorm.Engine
	var errDb error
	var objList []*models.DirInfo = make([]*models.DirInfo,0)

	db, errDb =  xorm.NewEngine("mysql", *mysqlAddr)
	if errDb != nil {
		panic(errDb)
	}
	db.DB().SetConnMaxLifetime(2000)
	exist , _:=db.IsTableExist(models.Dir_Inf{})
	if !exist{
		db.CreateTables(models.DirInfo{})
		db.CreateUniques(models.DirInfo{})
	}

	defer db.Close()
	count:=0
	for{
		select {
		case <- time.After(5 * time.Second):
			if count>0{
				fmt.Printf("ending with number %d \n",count)
				db.Insert(&objList)
			}
			return
		case obj:=<-objChannel:
			objList = append(objList,obj)
			count++
			if count ==150{
				fmt.Println(len(objList))
				db.Insert(&objList)
				count=0
				objList = make([]*models.DirInfo,0)
			}
		}
		//db.Exec("INSERT INTO dir_infs (path,is_dir,length,mod_time) VALUES ( ?,?,?,?);",nn.Path,nn.IsDir,nn.Length,nn.ModTime,"nilnil")
	}
}

func singlePersistence(){
	var db *xorm.Engine
	var errDb error
	selfId:= rand.Int()
	count:=0

	db, errDb =  xorm.NewEngine("mysql", *mysqlAddr)
	if errDb != nil {
		panic(errDb)
	}
	db.DB().SetConnMaxLifetime(2000)
	exist , _:=db.IsTableExist(models.Dir_Inf{})
	if !exist{
		db.CreateTables(models.DirInfo{})
		db.CreateUniques(models.DirInfo{})
	}

	defer db.Close()
	fmt.Printf("Id:%d start time:%s\n",selfId,time.Now())

	for{
		select {
		case <- time.After(5 * time.Second):
			fmt.Printf("Id:%d start time:%s,ending with number %d\n",selfId,time.Now(),count)
			endChannel<-count
			return
		case obj:=<-objChannel:
			_,err:=db.Insert(obj)
			if err!=nil{
				fmt.Println(err)
			}else{
				count++
			}
		}
	}
}

func main()  {
	flag.Parse()

	objChannel=make(chan *models.DirInfo,30*(*sqlConnNum))
	endChannel=make(chan int)

	for i:=0;i<*sqlConnNum;i++{
		go singlePersistence()
	}
	clients = make(chan *hdfs.Client, *clientNum)
	for i:=0;i<*clientNum;i++{
		client, err := hdfs.New(*hdfsDest)
		if err!=nil{
			panic(err)
		}
		clients<-client
	}

	go traverseDir(*rootDir)
	finishCount:=0
	totalItem:=0
	for{
		singleFinish:=<-endChannel
		totalItem+=singleFinish
		finishCount++
		if finishCount == *sqlConnNum{
			fmt.Printf("totally number:%d\n",totalItem)
			return
		}
	}
}

func traverseDir(input string){
	client :=<-clients
	fInfo,err:=  client.ReadDir(input)
	if err!=nil{
		panic(err)
	}
	clients<-client
	for _,v := range fInfo{
		//fmt.Println(input + v.Name())
		nn:=models.NewDir(v,input)
		objChannel<-nn
		if (v.IsDir()){
			buff:=bytes.Buffer{}
			buff.WriteString(input)
			buff.WriteString(v.Name())
			buff.WriteString("/")

			go traverseDir(buff.String())
		}
	}
}
