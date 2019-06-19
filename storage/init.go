package storage

import(
	"database/sql"
	"fmt"
	// "encoding/json"
	// "io/ioutil"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	config "overseas-bulter-server/config"
)

var db *sql.DB

type Config struct {
	dbHostIp string
	dbUserName string
	dbPassword string
	dbName string
}

var (
	devConfig = &Config{
		"0.0.0.0:3306",
		"root",
		"scwy1986614",
		"overseas_bulter"}
	proConfig = &Config{
		"0.0.0.0:3781",
		"root",
		"llyscysykr#WO&3OuEsnnes36$99324",
		"overseas_bulter"}
)

func createTables() {
	CreateUserTable()
	CreateHouseTable()
	CreateOrderTable()
	CreateFeedbackTable()
	CreateServiceTable()
	CreateEmployeeTable()
	CreateWechatTable()
	CreateBillingRecordTable()
	CreateInspectRecordTable()
	CreateRepairRecordTable()
	CreateServicePrimaryTable()
	CreateLayoutTable()
	CreatePriceParamsTable()
	CreateCarouselFigureTable()
	CreateHouseDealTable()
}

func getConfig() *Config{
	mode := config.RunMode
	databaseIp := config.DatabaseIp
	if("dev" == mode){
		devConfig.dbHostIp = databaseIp
		return devConfig
	}else if("pro" == mode){
		proConfig.dbHostIp = databaseIp
		return proConfig
	}else{
		fmt.Println("invalid mode in config json")
		return nil
	}
	
}

func Init() {
	config := getConfig();
	localDb,err := sql.Open("mysql",config.dbUserName+":"+config.dbPassword+"@tcp("+config.dbHostIp+")/"+config.dbName+"?charset=utf8")
	Error.CheckErr(err)
	fmt.Println("database open !");
	db = localDb
	createTables()
}