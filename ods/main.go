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
		"dbbos_latelygoodslist_parquet",
		"sapbwp_bic_azodw016300",
		"ration",
		"rationitem",
		"receipt",
		"receiptitem",
		"ret",
		"retitem",
		"transfer",
		"transferitem",
		"dim_bravo_shop",
		"fresh_must_have_list",
		"ods_union_customized_goods",
		"test20181024",
		"logistics_hc_transation_temp2_ekpo_ekko",
		"dim_wl_werks",
		"logistics_location_use_item",
		"logistics_pickfaceproduct_stationaspect",
		"test_goods_zp",
		"tmp",
		"tmp2",
		"dc_tmp5",
		"dbbos_ration_tmp",
		"sdbmgr_out_fcst",
		"b2b_ordercenter_prod_mch_audit_order",
		"b2b_ordercenter_prod_mch_audit_order_his",
		"b2b_ordercenter_prod_order_trace",
		"b2b_ordercenter_prod_order_trace_detail",
		"b2b_ordercenter_prod_shop_rep_cart_ai",
		"b2b_productcenter_base_product_shop",
		"masterdata_prod_md_replenish",
		"masterdata_prod_md_replenishment_parameter",
		"order_prod_po_merchant_order_detail",
		"order_prod_po_trace_order_detail",
		"report_prod_b2b_bu_report_jda_detail",
		"report_prod_vr_venderid_mapping",
		"report_prod_zb2b_cm_goods",
		"vss_report_supplier_employee_authority",
		"vss_vender_business",
		"vss_vender_business_authority",
		"vss_vender_business_collect",
		"yyxxb_vr_relation",
		"fjyh_card_balance",
		"fjyh_im_txn_header",
		"fjyh_im_txn_order",
		"fjyh_im_txn_order_checking",
		"posdm_center__posdm_order_header",
		"posdm_center__posdm_order_item",
		"posdm_center__posdm_order_item_promotion",
		"posdm_center__posdm_order_payment",
		"settlement_supplier_payable_center__promotion_chargeback_detail",
		"settlement_supplier_payable_center__supplier_bill",
		"ecc_eina",
		"ecc_eine",
		"ecc_ekko",
		"ecc_kna1",
		"ecc_knb1",
		"ecc_knbk",
		"ecc_knvp",
		"ecc_knvv",
		"ecc_lfa1",
		"ecc_lfb1",
		"ecc_lfbk",
		"ecc_lfm1",
		"ecc_likp",
		"ecc_lips",
		"ecc_mara",
		"ecc_marc",
		"ecc_mast",
		"ecc_mbew",
		"ecc_stko",
		"ecc_stpo",
		"ecc_t001l",
		"ecc_t001w",
		"ecc_t006",
		"ecc_t024",
		"ecc_t024e",
		"ecc_t171t",
		"ecc_t179",
		"ecc_tvko",
		"ecc_wrs1",
		"ecc_wyt3",
		"ecc_yt001w",
		"ecc_yt002_b",
		"ecc_ytactivehis",
		"ecc_ytdlcg",
		"ecc_ytdqexclude",
		"ecc_ytdqxzgz",
		"ecc_yteord",
		"ecc_ytfjcf",
		"ecc_ytgxupi",
		"ecc_ytjjdq",
		"ecc_ytlbcj",
		"ecc_ytloclb_cfg",
		"ecc_ytmadc",
		"ecc_ytmadq",
		"ecc_ytmdxzmx",
		"ecc_ytmdxzs",
		"ecc_ytmprice",
		"ecc_ytplgz",
		"ecc_ytpugc",
		"ecc_ytrort",
		"ecc_ytsfpnr",
		"ecc_ytstatu",
		"ecc_ytstatuhis",
		"ecc_ytstatuhis_jiti_quota",
		"ecc_ytstatuhis_jiti_quota_1",
		"sys_bic_yhbw_attribute_plant_att_time_dimension",
		"sapbwp_bi0_pcust_sales",
		"sapbwp_bi0_pcustomer",
		"sapbwp_bi0_pplant",
		"sapbwp_bi0_prt_asort",
		"sapbwp_bi0_prt_promo",
		"sapbwp_bic_azoco005500",
		"sapbwp_bic_azodw020300",
		"sapbwp_bic_azoma010400",
		"sapbwp_bic_azoma010700",
		"sapbwp_bic_azomp100000",
		"sapbwp_bic_azoop002b00",
		"sapbwp_bic_azoop005500",
		"sapbwp_bic_azopur04200",
		"sapbwp_bic_azopur95500",
		"sapbwp_bic_azopur91900",
		"sapbwp_bic_azopur94500",
		"sapbwp_bic_azuop001400",
		"sapbwp_bic_azwop900900",
		"sapbwp_bic_azwopb00000",
		"sapbwp_bic_azwop150100",
		"sapbwp_bic_azwop150200",
		"sapbwp_bi0_trt_promoct",
		"dim_sap_goods",
		"dim_sap_shop",
		"ecc_eban",
		"ecc_ekbe",
		"ecc_eket",
		"ecc_ekpo",
		"ecc_vbak",
		"ecc_yt218_a_all",
		"ecc_yt218_b_all",
		"sapecc_ytbms174l",
		"ecc_ytdzdj",
		"ecc_ytdzhd",
		"ecc_ytmeord",
		"hepecc_makt",
		"hepecc_mkpf",
		"hep_mseg",
		"hep_ytdzjy",
		"hepecc_ekko",
		"hepecc_t006a",
		"hepecc_vbrk",
		"hepecc_ytbms174l",
		"hepecc_ytdzhd",
		"hepecc_zvw_material",
		"hepecc_zvw_plant",
		"hepecc_zvw_pur_org_vendor",
		"inv_sap_setl_dly",
		"cal_ytbms186",
		"sapbwp_bic_azopur92700",
		"sapbwp_bic_azopur94400",
		"sapbwp_bic_azoihr00300",
		"sapecc_ekbe",
		"sapecc_eket",
		"sapecc_ekko",
		"sapecc_ekpo",
		"sapecc_mkpf",
		"sapecc_mseg",
		"sapecc_vbrk",
		"sapecc_vbrp",
		"sapecc_yt218_a",
		"sapecc_yt218_b",
		"ecc_ytbms174h",
		"ecc_ytbms174l",
		"sapecc_ytdhzq",
		"sapecc_ytloclb_cfg",
		"sapbecc_ytmarket",
		"yhbw_attribute_vendor",
		"sapbwp_ytbs174_if",
		"ytbw_co_config",
		"ecc_ytvkbur",
		"ecc_ytdzjy",
		"zft_catg_m_backend_income_result",
		"yhbw_op_zwopb000_ana_zwopb000",
		"address",
		"application_user",
		"inventory",
		"inventory_transaction",
		"invtr_sheettype",
		"location",
		"move_task",
		"order_header",
		"order_line",
		"pick_face",
		"pre_advice_header",
		"pre_advice_line",
		"shipping_manifest",
		"sku",
		"sku_config",
		"sku_sku_config",
		"supplier_sku",
		"upi_receipt_header",
		"upi_receipt_line",
		"cal_ychr_gs",
		"cal_ychr_lbr",
		"cal_ychr_zzcj",
		"cal_ychr_zzxx",
		"dbbos_cashiersheetdata",
		"k_results",
		"dim_member_feature",
		"t_class_group_credit",
		"t_gift",
		"t_gift_exchange",
		"t_gift_exchange_record",
		"t_gift_stock",
		"t_gift_stock_adjust",
		"t_gift_stock_allot",
		"t_member",
		"t_member_address",
		"t_member_card",
		"t_member_channel",
		"t_member_mobile",
		"t_member_mobile_9b",
		"t_member_mobile_not_9b",
		"t_member_mobile_original",
		"t_member_new",
		"t_promotion_activity_budget",
		"t_promotion_coupon_pool",
		"t_promotion_plan_sub_activity",
		"yh_credit_account",
		"yh_credit_adjust_bill",
		"yh_credit_cost_bill",
		"yh_credit_earn_bill",
		"yh_credit_event",
		"yh_credit_exchange_gift",
		"yh_credit_exchange_order",
		"yh_customer_request",
		"yh_member",
		"dbbos_appretgoodsitem",
		"dbbos_appretsheet",
		"dbbos_appsalesheet",
		"dbbos_appsalesheetitem",
		"dbbos_askcompensationlist",
		"dbbos_balanceplu",
		"dbbos_changecatesheet",
		"dbbos_changecatesheetacc",
		"dbbos_changeunitgoodsshop",
		"dbbos_combgoods",
		"dbbos_combgoodsitem",
		"dbbos_cstktakepl",
		"dbbos_cstktakeplitem",
		"dbbos_dayshopstock_sh",
		"dbbos_freshcost",
		"dbbos_goodscombine",
		"dbbos_goodsshop",
		"dbbos_inventorybook",
		"dbbos_latelygoodslist",
		"dbbos_lost",
		"dbbos_lostacc",
		"dbbos_materialdept",
		"dbbos_materialuse",
		"dbbos_materialuseacc",
		"dbbos_materialuseitem",
		"dbbos_orderadbook",
		"dbbos_orderadvice",
		"dbbos_orderaideitem",
		"dbbos_pay_j",
		"dbbos_placemove",
		"dbbos_placemoveitem",
		"dbbos_promshare",
		"dbbos_promshareitem",
		"dbbos_purchase",
		"dbbos_purchasechk",
		"dbbos_purchasechkitem",
		"dbbos_purchaseitem",
		"dbbos_ration",
		"dbbos_rationacc",
		"dbbos_rationitem",
		"dbbos_rconreceipt",
		"dbbos_rconreceiptitem",
		"dbbos_receipt",
		"dbbos_receiptitem",
		"dbbos_ret",
		"dbbos_retacc",
		"dbbos_retitem",
		"dbbos_rpgoodsstatus",
		"dbbos_rpt_clerkrate",
		"dbbos_rpt_freshskudaily_sh",
		"dbbos_sale_combine",
		"dbbos_sale_j",
		"dbbos_sale_pop_memo",
		"dbbos_salecard",
		"dbbos_salecardbalance",
		"dbbos_salecardbalancecancel",
		"dbbos_salecardcancel",
		"dbbos_salecarditem",
		"dbbos_splitgoodslist",
		"dbbos_splitsheet",
		"dbbos_splitsheetitem",
		"dbbos_stktakepl",
		"dbbos_stktakeplacc",
		"dbbos_stktakeplitem",
		"dbbos_stkvalueadj",
		"dbbos_stkvalueadjitem",
		"dbbos_transfer",
		"dbbos_transfergoods",
		"dbbos_transfergoodsitem",
		"dbbos_transferitem",
		"dbbos_wholeretgoodsitem",
		"dbbos_wholeretsheet",
		"dbbos_wholesalegoodsitem",
		"dbbos_wholesalesheet",
		"dbbos_guest_yyxx_customersheet_group",
		"dbbos_yyxx_salecostitem",
		"dbbos_guest_yyxx_salecostitem_group",
		"sharegoo_k_results",
		"dw_bravo_group_area_budget",
		"dw_bravo_group_budget",
		"dw_bravo_shops_budget",
		"d_goods_indate",
		"platform_d_grouplist",
		"platform_d_grouplist_sm",
		"dim_category_fresh_tree",
		"ods_fireports",
		"saima_shop_group_type",
	}
}

func main(){
	flag.Parse()
	for _,v := range tableList{
		//rootWg :=sync.WaitGroup{}
		//rootWg.Add(1)
		paths:= getPathsByName(v)

		for _,vv:=range paths{
			fmt.Println("get path:"+vv)
			checker,err := NewTableChecker(vv)
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
}

func getPathsByName(tableName string) []string{
	fmt.Println("begin check:"+tableName)
	engine, error := xorm.NewEngine("mysql", "root:metadata@Tbds.com@tcp(10.216.126.151:3306)/hive_back?charset=utf8")
	defer engine.Close()
	if error!=nil{
		return nil
	}

	//19935是代表DW这个db的地址,ODS 是165643
	//tlbs:= &models.Tbls{TblName:tableName}

	tlbs := make([]*models.Tbls, 0)
	err := engine.Where("TBL_NAME = ?",tableName).Find(&tlbs)
	//has ,err :=engine.Get(tlbs)
	//if (!has){
	//	return nil,fmt.Errorf("%s not exists1",tableName)
	//}
	if(err!=nil){
		fmt.Println(err)
		return nil
	}
	//fmt.Println(tlbs.SdId)
	result := make([]string,0)
	for _,v := range tlbs{
		sdid := &models.Sds{SdId:v.SdId}
		has,err := engine.Get(sdid)
		if (!has){
			continue
		}
		if(err!=nil){
			continue
		}

		if(strings.Index(sdid.Location,"hdfs://hdfsCluster")<0) || (strings.Index(sdid.Location,"warehouse")<0){
			continue
		}
		result =  append(result,sdid.Location[len("hdfs://hdfsCluster"):])
	}
	return result
}

type tableChecker struct {
	effectivePath string
	//updateList []*SyncItemOds
	//deleteList []*SyncItemOds
	//doublecheckList []*SyncItemOds
	finishChan chan struct{}
	SyncItemOdsChan chan *models.SyncItemOds

	SyncItemOdsSlice []*models.SyncItemOds
	SyncItemOdsMap map[string]*models.SyncItemOds

	writeEngine *xorm.Engine
	tbdsClients chan *hdfs.Client
	uatClients chan *hdfs.Client
	startTime time.Time
}
//后台一个goroutine来接收SyncItemOds
func (t *tableChecker) acceptItem(){
	for{
		select {
		case <-t.finishChan:
			break
		case SyncItemOdsTemp:=<-t.SyncItemOdsChan:
			if existingItem,exist:=t.SyncItemOdsMap[SyncItemOdsTemp.Path];!exist{
				t.SyncItemOdsMap[SyncItemOdsTemp.Path]= SyncItemOdsTemp
				t.SyncItemOdsSlice =append(t.SyncItemOdsSlice,SyncItemOdsTemp)
			}else{
				fmt.Printf("different %s,%s\n",SyncItemOdsTemp,existingItem)
			}
		}
	}
}

//接受各种各样的SyncItemOds
func (t *tableChecker) addDelete(input string){
	t.SyncItemOdsChan<-&models.SyncItemOds{
		Path:input,
		StartTime:t.startTime,
		Stage:0,
		MissionType:delete,
	}
}
func (t *tableChecker) addUpdate(input string){
	t.SyncItemOdsChan<-&models.SyncItemOds{
		Path:input,
		StartTime:t.startTime,
		Stage:0,
		MissionType:update,
	}
}
func (t *tableChecker) addUnexpect(input string){
	t.SyncItemOdsChan<- &models.SyncItemOds{
		Path:input,
		StartTime:t.startTime,
		Stage:0,
		MissionType:unexpected,
	}
}

//生成一个Checker
func NewTableChecker(effectivePath string) (*tableChecker,error){
	engineRecord, error := xorm.NewEngine("mysql", "root:DataLake_Yonghui1@tcp(10.216.155.15:3306)/migration?charset=utf8")
	if error!=nil{
		return nil,fmt.Errorf("%s create db conn error:%s",effectivePath,error)
	}

	fmt.Println(effectivePath)
	exist , _:=engineRecord.IsTableExist(models.SyncItemOds{})
	if !exist{
		error :=engineRecord.CreateTables(models.SyncItemOds{})
		if error!=nil{
			return nil,fmt.Errorf("%s create table error:%s",effectivePath,error)
		}
		//error =engineRecord.CreateUniques(SyncItemOds{})
		//if error!=nil{
		//	return nil,fmt.Errorf("%s create unique constraint error:%s",tableName,error)
		//}
	}

	tClients := make(chan *hdfs.Client, *clientNum)
	for i:=0;i<*clientNum;i++{
		client, err := hdfs.New(*tbdsDest)
		if err!=nil{
			return nil,fmt.Errorf("%s get tbds client fail:%s",effectivePath,error)
		}
		tClients<-client
	}
	uClients := make(chan *hdfs.Client, *clientNum)
	for i:=0;i<*clientNum;i++{
		client, err := hdfs.New(*hdfsDest)
		if err!=nil{
			return nil,fmt.Errorf("%s get uat client fail:%s",effectivePath,error)
		}
		uClients<-client
	}

	return &tableChecker{
		effectivePath:effectivePath,
		SyncItemOdsChan: make(chan *models.SyncItemOds,200),
		//rootWg:sync.WaitGroup{},
		finishChan:make(chan struct{}),

		SyncItemOdsSlice: make([]*models.SyncItemOds,0) ,
		SyncItemOdsMap:make(map[string]*models.SyncItemOds),

		writeEngine:engineRecord,
		tbdsClients:tClients,
		uatClients:uClients,
		startTime:time.Now(),
	},nil
}

func (t *tableChecker)persistent() error{
	//todo：这里的同步机制需要更好的设计
	time.Sleep(3*time.Second)
	//停止掉接受SyncItemOds的协程
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

	for len(t.SyncItemOdsSlice)>200{
		inserting := make([]*models.SyncItemOds,200)
		copy(inserting,t.SyncItemOdsSlice[:200])
		_,err:=t.writeEngine.Insert(&inserting)
		if err!=nil{
			return fmt.Errorf("%s insert err: %s",t.effectivePath,err)
		}
		t.SyncItemOdsSlice = t.SyncItemOdsSlice[200:]
	}
	if len(t.SyncItemOdsSlice)!=0{
		_,err:=t.writeEngine.Insert(&t.SyncItemOdsSlice)
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
	inited := false
	var isDir bool
	for i,v :=range input{
		//跳过.schema的目录
		if strings.Index(v.Name(),".schemas")>=0{
			continue
		}
		if(!inited){
			isDir=v.IsDir()
			inited = true
		}
		if isDir!=v.IsDir(){
			fmt.Printf("********:%s:%d:%s\n",v.Name(),i,input[0].Name())
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