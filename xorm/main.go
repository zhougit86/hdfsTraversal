package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"yonghui.com/hdfsTraver/hdfs/models"
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"flag"
	"strings"
	"bytes"
	"os"
	"time"
)

var tableList []string

//var engine *xorm.Engine
//var engineRecord *xorm.Engine

var nowTime time.Time
var hdfsDest *string
var tbdsDest *string
var clientNum *int

const (
	fileDir = iota
	dirDir
	InvalidDir
	emptyDir
)

const (
	update = iota
	delete
	unexpected
)

func init()  {
	hdfsDest= flag.String("hdfs", "10.1.53.205:8020", "The hdfs you want to connect")
	tbdsDest= flag.String("tbds", "10.216.126.151:8020", "The tbds you want to connect")
	clientNum = flag.Int("client",3,"the number of client to hdfs")

	nowTime = time.Now()
	tableList = []string{
		"catg_change_dtl_fct",
		"catg_goods_active_change_fct",
		"catg_goods_change_fct",
		"catg_goods_rp_flag_fct",
		"catg_goods_statu_change_fct",
		"catg_logistics_wstg_fct",
		"catg_main_goods_trans_fct",
		"catg_organic_wstg_fct",
		"catg_sku_askcompensationlist_fct",
		"catg_sku_loss_fct",
		"catg_sku_stktakepl_fct",
		"catg_sku_stkvalueadj_fct",
		"catg_sku_wstg_fct",
		"catg_split_rlt_fct",
		"fin_acct_g2_dept_mon",
		"fin_card_balance_chloy_fct",
		"fin_card_balance_fct",
		"fin_card_ord_txn_chloy_fct",
		"fin_card_ord_txn_head_chloy_fct",
		"fin_card_pay_fct",
		"fin_card_pay_sales_share_fct",
		"fin_dept_gross_margin_fct",
		"fin_dept_race_kpi_mtd",
		"fin_front_gross_margin_fct",
		"fin_pro_chg_item_doc_fct",
		"fin_pro_chg_item_fct",
		"fin_race_base_mon",
		"fin_race_base_mtd",
		"fin_statement_head_fct",
		"fin_statement_item_fct",
		"fin_vendor_backend_income_fct",
		"fin_vendor_catg_m_backend_income_fct",
		"hep_fin_statement_head_fct",
		"hrm_attendance_fct",
		"hrm_cashier_quality_fct",
		"hrm_clerk_efficient_rpt",
		"hrm_cs_attendance_fct",
		"hrm_cs_cashier_stat_fct",
		"hrm_cs_cosurety_fct",
		"hrm_cs_horgunit_fct",
		"hrm_cs_work_hrs_fct",
		"hrm_person_info_fct",
		"inv_class_dos_fct",
		"inv_detail_account_fct",
		"inv_dtl_acct_goods_latest",
		"inv_goods_change_summary_mon",
		"inv_goods_dos_fct",
		"inv_hour_fct",
		"inv_onway_dly_fct",
		"inv_sap_setl_dly_fct",
		"inv_setl_dly_fct",
		"inv_setl_dly_fct_backup",
		"inv_shop_goods_remain_dly",
		"inv_stat_mon",
		"inv_tax_amt",
		"inv_transfer_fct",
		//"member_feature_belong_shop_fct",
		//"member_feature_consume_analysis",
		//"member_feature_first_shop_fct",
		//"member_feature_goods_analysis",
		//"member_feature_recent_shop_fct",
		//"member_feature_shop_analysis",
		//"member_feature_stage_analysis",
		//"member_feature_time_period_analysis",
		"mkt_pro_channel_fct",
		"mkt_pro_evt_cust_fct",
		"mkt_pro_evt_detail_fct",
		"mkt_pro_evt_head_fct",
		"mkt_pro_evt_rule_detail_fct",
		"mkt_pro_evt_sch_fct",
		"mkt_pro_rule_detail_fct",
		"mkt_pro_rule_fct",
		"mkt_pro_rule_head_fct",
		"mkt_pro_shop_fct",
		"mkt_shop_goods_market_flag",
		"sale_backgm_fct",
		"sale_backgm_mon",
		"sale_city_catg_nielsen_wek",
		"sale_city_goods_nielsen_mon",
		"sale_comb_goods_stat",
		"sale_dms_fct",
		"sale_groupon_fct",
		"sale_groupon_ret_fct",
		"sale_new_goods_fct",
		"sale_prov_catg_nielsen_mon",
		"sale_rt_receipt_fct",
		"sale_sap_dtl_fct",
		"sale_sap_order_fct",
		"sale_setl_dly_fct",
		"sale_shopping_card_fct",
		"sale_shopping_card_item_fct",
		"sale_stat_dly",
		"sale_stat_mon",
		"shop_all_custflow_dly",
		"shop_bd_custflow_hour",
		"shop_catg_l_custflow_hour",
		"shop_catg_m_custflow_hour",
		"shop_catg_s_custflow_hour",
		"shop_code_scan_mon",
		"shop_custflow_bd_custflow",
		"shop_custflow_bd_custflow_hour",
		"shop_custflow_bd_custflow_hour_update",
		"shop_custflow_hour",
		"shop_dept_custflow_hour",
		"shop_dept_custflow_hour_test",
		"shop_dept_custflow_hour_time",
		"shop_dept_custflow_hour_time_update",
		"shop_dept_custflow_hour_time_update_test",
		"shop_div_custflow_hour",
		"shop_firm_bd_custflow_hour",
		"shop_firm_dept_skus_hour",
		"shop_firm_g1_custflow_hour",
		"shop_firm_g2_custflow_hour",
		"shop_goods_custflow_dly",
		"shop_goods_custflow_mon",
		"shop_goods_fct",
		"shop_goods_fct_main_goods",
		"shop_new_bd_custflow_dly",
		"shop_sap_custflow_fct",
		"supp_cop_delivery_fct",
		"supp_dc_order_fct",
		"supp_delivery_amt_fct",
		"supp_delivery_fct",
		"supp_inv_fct",
		"supp_inv_item",
		"supp_inv_item_return",
		"supp_inv_sheet_type",
		"supp_inv_transaction_fct",
		"supp_inv_transaction_fct_item",
		"supp_inv_txn_action_type",
		"supp_material_use_fct",
		"supp_move_task_fct",
		"supp_online_ret_dtl_fct",
		"supp_online_sale_dtl_fct",
		"supp_pack_receipt_fct",
		"supp_pur_check_fct",
		"supp_pur_fct",
		"supp_pur_sale_sts_fct",
		"supp_receipt_fct",
		"supp_receipt_ret_fct",
		"supp_sap_delivery_fct",
		"supp_sap_pur_order_fct",
		"supp_sap_pur_order_his_fct",
		"supp_sap_pur_plan_fct",
		"supp_sap_vendor_pur_ret_fct",
		"supp_transfer_fct",
		"supp_upi_receipt_fct",
	}
}

func main(){
	flag.Parse()
	for _,v := range tableList{
		//rootWg :=sync.WaitGroup{}
		//rootWg.Add(1)
		checker,err := NewTableChecker(v)
		if err!=nil{
			fmt.Println(err)
			continue
		}
		go checker.acceptItem()
		checker.traverseDir(checker.effectivePath+"/")
		err = checker.persistent()
		if err!=nil{
			fmt.Println(err)
		}
	}
}

type tableChecker struct {
	effectivePath string
	//updateList []*SyncItem
	//deleteList []*SyncItem
	//doublecheckList []*SyncItem
	finishChan chan struct{}
	syncItemChan chan *models.SyncItem

	syncItemSlice []*models.SyncItem
	syncItemMap map[string]*models.SyncItem

	writeEngine *xorm.Engine
	tbdsClients chan *hdfs.Client
	uatClients chan *hdfs.Client
	startTime time.Time
}
//后台一个goroutine来接收SyncItem
func (t *tableChecker) acceptItem(){
	for{
		select {
		case <-t.finishChan:
			break
		case syncItemTemp:=<-t.syncItemChan:
			if existingItem,exist:=t.syncItemMap[syncItemTemp.Path];!exist{
				t.syncItemMap[syncItemTemp.Path]= syncItemTemp
				t.syncItemSlice =append(t.syncItemSlice,syncItemTemp)
			}else{
				fmt.Printf("different %s,%s\n",syncItemTemp,existingItem)
			}
		}
	}
}

//接受各种各样的SyncItem
func (t *tableChecker) addDelete(input string){
	t.syncItemChan<-&models.SyncItem{
		Path:input,
		StartTime:t.startTime,
		Stage:0,
		MissionType:delete,
	}
}
func (t *tableChecker) addUpdate(input string){
	t.syncItemChan<-&models.SyncItem{
		Path:input,
		StartTime:t.startTime,
		Stage:0,
		MissionType:update,
	}
}
func (t *tableChecker) addUnexpect(input string){
	t.syncItemChan<- &models.SyncItem{
		Path:input,
		StartTime:t.startTime,
		Stage:0,
		MissionType:unexpected,
	}
}

//生成一个Checker
func NewTableChecker(tableName string) (*tableChecker,error){
	fmt.Println("begin check:"+tableName)
	engine, error := xorm.NewEngine("mysql", "root:metadata@Tbds.com@tcp(10.216.126.151:3306)/hive_back?charset=utf8")
	defer engine.Close()
	if error!=nil{
		return nil,fmt.Errorf("%s got error:%s",tableName,error)
	}

	//19935是代表DW这个db的地址
	tlbs:= &models.Tbls{TblName:tableName,DbId:19935}
	has ,err :=engine.Get(tlbs)
	if (!has){
		return nil,fmt.Errorf("%s not exists1",tableName)
	}
	if(err!=nil){
		return nil,fmt.Errorf("%s not exists2",tableName)
	}
	//fmt.Println(tlbs.SdId)
	sdid := &models.Sds{SdId:tlbs.SdId}
	has,err = engine.Get(sdid)
	if (!has){
		return nil,fmt.Errorf("%s not exists3",tableName)
	}
	if(err!=nil){
		return nil,fmt.Errorf("%s not exists4",tableName)
	}

	if(strings.Index(sdid.Location,"hdfs://hdfsCluster")<0){
		return nil,fmt.Errorf("%s can not parse path:%s",tableName,sdid.Location)
	}
	effectivePath := sdid.Location[len("hdfs://hdfsCluster"):]

	engineRecord, error := xorm.NewEngine("mysql", "root:DataLake_Yonghui1@tcp(10.216.155.15:3306)/migration?charset=utf8")
	if error!=nil{
		return nil,fmt.Errorf("%s create db conn error:%s",tableName,error)
	}

	fmt.Println(effectivePath)
	exist , _:=engineRecord.IsTableExist(models.SyncItem{})
	if !exist{
		error :=engineRecord.CreateTables(models.SyncItem{})
		if error!=nil{
			return nil,fmt.Errorf("%s create table error:%s",tableName,error)
		}
		//error =engineRecord.CreateUniques(SyncItem{})
		//if error!=nil{
		//	return nil,fmt.Errorf("%s create unique constraint error:%s",tableName,error)
		//}
	}

	tClients := make(chan *hdfs.Client, *clientNum)
	for i:=0;i<*clientNum;i++{
		client, err := hdfs.New(*tbdsDest)
		if err!=nil{
			return nil,fmt.Errorf("%s get tbds client fail:%s",tableName,error)
		}
		tClients<-client
	}
	uClients := make(chan *hdfs.Client, *clientNum)
	for i:=0;i<*clientNum;i++{
		client, err := hdfs.New(*hdfsDest)
		if err!=nil{
			return nil,fmt.Errorf("%s get uat client fail:%s",tableName,error)
		}
		uClients<-client
	}

	return &tableChecker{
		effectivePath:effectivePath,
		syncItemChan: make(chan *models.SyncItem,200),
		//rootWg:sync.WaitGroup{},
		finishChan:make(chan struct{}),

		syncItemSlice: make([]*models.SyncItem,0) ,
		syncItemMap:make(map[string]*models.SyncItem),

		writeEngine:engineRecord,
		tbdsClients:tClients,
		uatClients:uClients,
		startTime:time.Now(),
	},nil
}

func (t *tableChecker)persistent() error{
	//todo：这里的同步机制需要更好的设计
	time.Sleep(3*time.Second)
	//停止掉接受syncitem的协程
	t.finishChan<- struct{}{}
	defer t.writeEngine.Close()
	defer func() {
		for i:=0;i<3;i++{
			v:=<-t.tbdsClients
			v.Close()
		}
	}()
	defer func() {
		for i:=0;i<3;i++{
			v:=<-t.uatClients
			v.Close()
		}
	}()

	for len(t.syncItemSlice)>200{
		inserting := make([]*models.SyncItem,200)
		copy(inserting,t.syncItemSlice[:200])
		_,err:=t.writeEngine.Insert(&inserting)
		if err!=nil{
			return fmt.Errorf("%s insert err: %s",t.effectivePath,err)
		}
		t.syncItemSlice = t.syncItemSlice[200:]
	}
	if len(t.syncItemSlice)!=0{
		_,err:=t.writeEngine.Insert(&t.syncItemSlice)
		if err!=nil{
			return err
		}
	}

	//if(len(t.updateList)!=0){
	//	_,err:=engineRecord.Insert(&t.updateList)
	//	if err!=nil{
	//		return err
	//	}
	//}
	//if(len(t.deleteList)!=0){
	//	_,err:=engineRecord.Insert(&t.deleteList)
	//	if err!=nil{
	//		return err
	//	}
	//}
	//if(len(t.doublecheckList)!=0){
	//	_,err:=engineRecord.Insert(&t.doublecheckList)
	//	if err!=nil{
	//		return err
	//	}
	//}
	return nil
}

//根据文件的列表来获取一个map
func generateMap(input []os.FileInfo) (map[string]os.FileInfo,int){
	retMap := make(map[string]os.FileInfo)
	if len(input)==0{
		return retMap,emptyDir
	}

	var isDir bool
	for i,v :=range input{
		if strings.Index(v.Name(),".schemas")>=0{
			continue
		}
		if(i==0){
			isDir=v.IsDir()
		}
		if isDir!=v.IsDir(){
			return nil,InvalidDir
		}
		retMap[v.Name()] = v
	}
	if isDir{
		return retMap,dirDir
	}
	return retMap,fileDir
}

func (t *tableChecker) traverseDir(input string){

	uatClient:=<-t.uatClients
	UatFileInfo,err :=  uatClient.ReadDir(input)
	//Todo:判一下error的类型,事实上UAT上肯定有这个目录，不用判这个错
	if err!=nil{
		//如果在UAT没有则需要删除他
		t.addUnexpect(input)
		t.uatClients<-uatClient
		return
	}
	t.uatClients<-uatClient

	tbdsClient:=<-t.tbdsClients
	TbdsFileInfo,err :=  tbdsClient.ReadDir(input)
	//Todo:判一下error的类型
	if err!=nil{
		//如果在TBDS那边没有这个目录，就意味着需要从福州同步这个目录
		t.addUpdate(input)
		t.tbdsClients<-tbdsClient
		return
	}
	t.tbdsClients<-tbdsClient

	sourceMap ,sourceType := generateMap(UatFileInfo)
	if sourceType==InvalidDir{
		//todo:改回panic
		//panic("Uat The Dir has both type: "+input)
		fmt.Println("Uat The Dir has both type: "+input)
		t.addUnexpect(input)
		return
	}

	dstMap ,dstType := generateMap(TbdsFileInfo)
	if dstType==InvalidDir{
		//todo:改回panic
		//panic("Tbds The Dir has both type: "+input)
		fmt.Println("Tbds The Dir has both type: "+input)
		t.addUnexpect(input)
		return
	}else if dstType==emptyDir && sourceType!=emptyDir{
		t.addUpdate(input)
		return
	}

	//如果判下来两者的类型不同就需要报错
	if(dstType!=sourceType){
		t.addUnexpect(input)
	}

	//当是文件目录的时候，两边只要有不一致就把
	if sourceType ==fileDir{
		unIdentical := false
		//遍历源端的目录，当目标端缺少某个文件或者判定文件不一致，就要同步该父级目录
		for k,v :=range sourceMap{
			//fmt.Println(input + ":file:" +v.Name())
			dstCounterPart,ok:=dstMap[k]
			if !ok{
				t.addUpdate(input)
				unIdentical = true
				break
			}
			if !checkIdentical(v,dstCounterPart){
				t.addUpdate(input)
				unIdentical = true
				break
			}
		}
		//然后遍历目标端，如果源端没有的就删掉
		if !unIdentical{
			for k,v:=range dstMap{
				_,ok:=sourceMap[k]
				if !ok{
					buff:=bytes.Buffer{}
					buff.WriteString(input)
					buff.WriteString(v.Name())

					t.addDelete(buff.String())
				}
			}
		}
	}else{
		//如果是目录型的目录，源端有目标端没有，则同步该目录
		//两边都有则递归方法检查该目录
		//目标端有源端没有，jiu
		for k,v := range sourceMap{
			//fmt.Println(input + ":dir:" +v.Name())
			buff:=bytes.Buffer{}
			buff.WriteString(input)
			buff.WriteString(v.Name())
			buff.WriteString("/")

			_,ok:=dstMap[k]
			if !ok{
				//如果不存在就需要同步
				t.addUpdate(buff.String())
				continue
			}
			//如果存在就继续遍历
			t.traverseDir(buff.String())
		}
		for k,v:=range dstMap{
			buff:=bytes.Buffer{}
			buff.WriteString(input)
			buff.WriteString(v.Name())
			buff.WriteString("/")

			_,ok:=sourceMap[k]
			if !ok{
				//DeleteDir<-buff.String()
				t.addDelete(buff.String())
				continue
			}
		}
	}
}

func checkIdentical(src,dst os.FileInfo) bool{
	return src.Size()==dst.Size()
}