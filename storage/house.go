package storage

import(
	// "database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbHouse struct{
	Uid 		string
	Name 		string
	Country 	string
	Province 	string
	City 		string
	Address 	string
	Layout 		string
	OwnerId		string
	CreateTime	string
}

const(
	create_house_table_sql = `CREATE TABLE IF NOT EXISTS house_t(
		uid VARCHAR(64) NOT NULL unique,
		name VARCHAR(64) NULL DEFAULT NULL,
		country VARCHAR(64) NULL DEFAULT NULL,
		province VARCHAR(64) NULL DEFAULT NULL,
		city VARCHAR(64) NULL DEFAULT NULL,
		address VARCHAR(1024) NULL DEFAULT NULL,
		layout VARCHAR(64) NULL DEFAULT NULL,
		owner_id VARCHAR(64) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_house = `INSERT INTO house_t (uid,name,country,province,city,address,layout, owner_id ) value (?,?,?,?,?,?,?,?)`
	query_house_by_range = `SELECT * FROM house_t LIMIT ? OFFSET ?`
	query_house = `SELECT * FROM house_t WHERE uid=?`
	update_house_by_uid = `UPDATE house_t SET name=?,country=?,province=?,city=?,address=?,layout=?,owner_id=? WHERE uid=?`
	delete_house_by_uid = `DELETE FROM house_t WHERE uid=?`
)

func CreateHouseTable() {
	sql := create_house_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create house table!");
	// arr := QueryHousesByRange(5, 0)
}

func QueryHouse(id string) *DbHouse{
	
	result := &DbHouse{}
	rows, err := db.Query(query_house, id)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Name, 
			&result.Country,  
			&result.Province, 
			&result.City, 
			&result.Address, 
			&result.Layout,
			&result.OwnerId,
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr house :")
	fmt.Println(result)
	return result
}

func AddHouse(
	name string,
	country string,
	province string,
	city string,
	address string,
	layout string,
	ownerId string) string{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_house ,uuid, name, country, province, city, address, layout, ownerId);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert house success !")
	fmt.Println(aff_nums);
	return uuid;
}

func UpdateHouse(
	uid string,
	name string,
	country string,
	province string,
	city string,
	address string,
	layout string,
	ownerId string){
	
	ret, err := db.Exec(update_house_by_uid, 
		name, 
		country, 
		province, 
		city, 
		address, 
		layout, 
		ownerId, 
		uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update house success !")
	fmt.Println(aff_nums)
}

func QueryHouses(count uint, offset uint) []DbHouse{
	
	result := make([]DbHouse, 0)
	rows, err := db.Query(query_house_by_range, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbHouse{}
		err = rows.Scan(
			&item.Uid,
			&item.Name,
			&item.Country,
			&item.Province,
			&item.City,
			&item.Address,
			&item.Layout,
			&item.OwnerId,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func DeleteHouseByUid(uid string){
	//删除数据
	ret, err := db.Exec(delete_house_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete house success !")
	fmt.Println(aff_nums);
}