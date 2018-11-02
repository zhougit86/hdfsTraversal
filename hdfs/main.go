package main

import(
	"github.com/colinmarc/hdfs/v2"

	"flag"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yonghui.com/hdfsTraver/hdfs/models"
	//"fmt"
	"bytes"
	"fmt"
)

var clients chan *hdfs.Client
var objChannel chan *models.Dir_Inf
//var clients []*hdfs.Client
var hdfsDest *string
var rootDir *string
var mysqlAddr *string
var clientNum *int
var sqlConnNum *int
var db *gorm.DB

func init()  {
	mysqlAddr= flag.String("sql", "root:123@tcp(192.168.13.128:3306)/migration?charset=utf8", "sql:root:123@tcp(192.168.13.128:3306)/migration?charset=utf8")
	hdfsDest= flag.String("hdfs", "192.168.13.128:9000", "The hdfs you want to connect")
	rootDir = flag.String("root" ,"/","the directory to start with")
	clientNum = flag.Int("client",3,"the number of client to hdfs")
	sqlConnNum = flag.Int("sqlconn",10,"the number of sql conn")
}

func Persistence(){
	buff:=bytes.Buffer{}
	buff.WriteString("INSERT INTO dir_infs (path,is_dir,length,mod_time) VALUES ")
	for i:=0;i<4999;i++{
		buff.WriteString("( ?,?,?,?),")
	}
	buff.WriteString("( ?,?,?,?);")
	var errDb error
	var objList []interface{} = make([]interface{},5000*4)

	db, errDb = gorm.Open("mysql", *mysqlAddr)
	if errDb != nil {
		panic(errDb)
	}
	db.DB().SetConnMaxLifetime(2000)
	if !db.HasTable(&models.Dir_Inf{}){
		db.CreateTable(&models.Dir_Inf{})
	}

	defer db.Close()
	count:=0
	for{
		obj:=<-objChannel
		objList[count*4] = obj.Path
		objList[count*4+1] = obj.IsDir
		objList[count*4+2] = obj.Length
		objList[count*4+3] = obj.ModTime
		count++
		if count ==5000{
			fmt.Println(len(objList))
			db.Exec(buff.String(),objList...)
			count=0
			objList = make([]interface{},5000*4)
		}
		//db.Exec("INSERT INTO dir_infs (path,is_dir,length,mod_time) VALUES ( ?,?,?,?);",nn.Path,nn.IsDir,nn.Length,nn.ModTime,"nilnil")
	}
}

func main()  {
	flag.Parse()

	objChannel=make(chan *models.Dir_Inf,30*(*sqlConnNum))

	for i:=0;i<*sqlConnNum;i++{
		go Persistence()
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
	time.Sleep(100000*time.Second)
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
			go traverseDir(input + v.Name() + "/")
		}
	}
}
