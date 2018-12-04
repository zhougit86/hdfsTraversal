package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"io/ioutil"
	"time"
	"github.com/go-xorm/xorm"
	"os"
	"strconv"
)

var engine *xorm.Engine

const (
	url = "/rest/sync_item/"
	Accepted  = "accepted"
	Fail  = "fail"
	Running  = "running"
	Finish  = "finish"
)

func GoStdTime()string{
	return "2006-01-02 15:04:05"
}

func syncItem(p *ParsingForm) (int,error) {
	if p==nil{
		return 0,fmt.Errorf("nil p pointer")
	}

	//解析时间的属性
	loc ,_ := time.LoadLocation("Asia/Shanghai")
	receivedTime,err:=  time.ParseInLocation(GoStdTime(),p.Date,loc)
	if err!=nil{
		fmt.Println(err)
		return 0,fmt.Errorf("parse date fail")
	}

	//首先搜索是否存在记录
	groupExisting := &GroupRecord{GroupId: int64(p.GroupId)}
	has,err:=engine.Get(groupExisting)
	if err!=nil{
		fmt.Println(err)
		return 0,fmt.Errorf("query fail")
	}
	//fmt.Println(groupExisting.Date)
	//fmt.Println(receivedTime)
	if has && groupExisting.Date.Equal(receivedTime) {
		return 0,fmt.Errorf("the record is not updated")
	}

	//如果存在则删除再创建，不存在则创建
	_,err=engine.Exec("replace into group_record(groupId, date) values(?, ?);",p.GroupId,receivedTime)
	if err!=nil{
		fmt.Println(err)
		return 0,fmt.Errorf("replace group record fail")
	}

	jobs := make([]*JobRecord,len(p.Jobs))
	for i,v := range p.Jobs{
		job := &JobRecord{
			Location: v.Location,
			//Date:receivedTime,
			JobId:   int64(v.JobId),
			GroupId: int64(p.GroupId),
			Status:  Accepted,
		}
		jobs[i]= job
	}
	_,err=engine.Insert(&jobs)
	if err!=nil{
		fmt.Println(err)
		return 0,fmt.Errorf("insert job record fail")
	}
	return len(jobs), nil
}

func getGroupId(groupId int) (string,error){
	pEveryOne := make([]*JobRecord, 0)
	err := engine.Where("groupId = ?",groupId).Find(&pEveryOne)
	if err !=nil{
		return "",err
	}
	if len(pEveryOne)==0{
		return "",fmt.Errorf("no entry found in db")
	}
	successNum := 0
	for _,v := range pEveryOne{
		if v.Status == Fail{
			return Fail,nil
		}
		if v.Status == Finish{
			successNum++
		}
	}
	if successNum==len(pEveryOne){
		return Finish,nil
	}
	return Running,nil
}

type (
	Job struct {
		JobId int `json:"jobId"`
		Location string `json:"location"`
	}
	ParsingForm struct {
		GroupId  int    `json:"groupId"`
		Date string `json:"date"`
		Jobs []*Job `json:"jobs"`
	}
	JobRecord struct {
		Id int64 `xorm:"pk autoincr notnull"`
		Location string `xorm:"notnull unique(multiUni) 'location'"`
		//Date time.Time `xorm:"notnull unique(multiUni) 'date'"`
		JobId int64 `xorm:" 'jobId'"`
		GroupId int64 `xorm:"notnull unique(multiUni) 'groupId'"`
		Status string `xorm:" 'status'"`
	}
	GroupRecord struct {
		//Location string `xorm:"notnull unique(multiUni) 'location'"`
		Date time.Time `xorm:"notnull 'date'"`
		//JobId int64 `xorm:" 'jobId'"`
		GroupId int64 `xorm:"notnull pk unique(multiUni) 'groupId'"`
		//Status string `xorm:" 'status'"`
	}
)

func (s *JobRecord) String() string{
	return fmt.Sprintf("%s:%d",s.Location,s.GroupId)
}

func main() {
	//设置时区等信息
	loc ,_:= time.LoadLocation("Asia/Shanghai")
	var error error
	engine, error = xorm.NewEngine("mysql", "root:DataLake_Yonghui1@tcp(10.216.155.15:3306)/migration?charset=utf8&interpolateParams=true&parseTime=true&loc=Local")
	if error!=nil{
		fmt.Println(error)
		os.Exit(1)
	}
	engine.DatabaseTZ = loc
	engine.TZLocation = loc

	exist , _:=engine.IsTableExist(JobRecord{})
	if !exist{
		error :=engine.CreateTables(JobRecord{})
		if error!=nil{
			fmt.Println(error)
			os.Exit(1)
		}
		engine.CreateUniques(JobRecord{})
	}

	exist , _=engine.IsTableExist(GroupRecord{})
	if !exist{
		error :=engine.CreateTables(GroupRecord{})
		if error!=nil{
			fmt.Println(error)
			os.Exit(1)
		}
		engine.CreateUniques(GroupRecord{})
	}

	http.HandleFunc(url, route)
	ret := http.ListenAndServe(":8080", nil)
	fmt.Println(ret)
}

func route(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resource := strings.TrimLeft(r.URL.Path ,url)
		groupId,err := strconv.Atoi(resource)
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		result,err := getGroupId(groupId)
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if result==Fail{
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(result))
	case "POST":

		body, err := ioutil.ReadAll(r.Body)
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("get body error"))
			return
		}
		p := &ParsingForm{}
		err = json.Unmarshal(body, p)

		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("decode error"))
			return
		}

		sum ,_:=syncItem(p)
		strconv.Itoa(sum)
		w.Write([]byte(strconv.Itoa(sum)))
	}
}



