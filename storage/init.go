package storage

import(
	"database/sql"
	"fmt"
	// "encoding/json"
	// "io/ioutil"
	_"github.com/go-sql-driver/mysql"
	"github.com/akkuman/parseConfig"
	Error "overseas-bulter-server/error"
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
		"0.0.0.0:3306",
		"root",
		"root123",
		"overseas_bulter"}
)

func createTables() {
	CreateUserTable()
	CreateHouseTable()
	CreateCaseTable()
	CreateOrderTable()
	CreateFeedbackTable()
	CreateServiceTable()
	CreateEmployeeTable()
	CreateWechatTable()
	CreateRentRecordTable()
	CreateInspectRecordTable()
	CreateRepairRecordTable()
}

func getConfig() *Config{
	var config = parseConfig.New("config.json")
	var mode = config.Get("mode")
	fmt.Println("config mode = ", mode);
	if("dev" == mode){
		return devConfig
	}else if("pro" == mode){
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