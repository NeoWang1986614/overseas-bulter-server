package storage

import(
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbService struct{
	Uid 		string
	Type 		string
	Layout		string
	Content		string
	Price	 	uint
	CreateTime	string
}

const(
	create_service_table_sql = `CREATE TABLE IF NOT EXISTS service_t(
		uid VARCHAR(64) NOT NULL unique,
		type VARCHAR(64) NULL DEFAULT NULL,
		layout VARCHAR(64) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		price INT(64) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_service = `INSERT INTO service_t (uid, type, layout, content, price) value (?,?,?,?,?)`
	query_service = `SELECT * FROM service_t WHERE uid=?`
	query_services_by_type = `SELECT * FROM service_t WHERE type=?`
	update_service_by_id = `UPDATE service_t SET type=?, layout=?, content=?, price=? WHERE uid=?`
)

func CreateServiceTable() {
	sql := create_service_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create service table!");
}

func QueryService(id string) *DbService{
	
	result := &DbService{}
	rows, err := db.Query(query_service, id)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Type, 
			&result.Layout, 
			&result.Content,  
			&result.Price,  
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr order :")
	fmt.Println(result)
	return result
}

func AddService(
	serviceType 	string,
	layout			string,
	content 		string,
	price			uint) {
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_service ,uuid, serviceType, layout, content, price);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert service success !")
	fmt.Println(aff_nums);
}

func QueryServicesByType(serviceType string) []DbService{

	result := make([]DbService, 0)
	rows, err := db.Query(query_services_by_type, serviceType)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbService{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Layout,
			&item.Content,
			&item.Price,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func UpdateService(
	id				string,
	serviceType		string,
	layout			string,
	content 		string,
	price			uint) {
	//更新数据
	ret, err := db.Exec(update_service_by_id ,serviceType, layout, content, price ,id);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update service success !")
	fmt.Println(aff_nums);
}