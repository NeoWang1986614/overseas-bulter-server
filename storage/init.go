package storage

import(
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
)

var db *sql.DB

var (
	dbHostIp="0.0.0.0:3306"//"129.28.57.139:3306"//"localhost:3306"//"0.0.0.0:3306"//
	dbUserName="root"
	dbPassword="scwy1986614"//
	dbName="overseas_bulter"
)

func createTables() {
	CreateUserTable()
	CreateHouseTable()
	CreateCaseTable()
	CreateOrderTable()
	CreateFeedbackTable()
	CreateServiceTable()
}

func Init() {
	localDb,err := sql.Open("mysql",dbUserName+":"+dbPassword+"@tcp("+dbHostIp+")/"+dbName+"?charset=utf8")
	Error.CheckErr(err)
	fmt.Println("database open !");
	db = localDb
	createTables()
}