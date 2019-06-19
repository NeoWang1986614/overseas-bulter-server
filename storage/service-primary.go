package storage

import(
	"fmt"
	"errors"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbServicePrimary struct{
	Uid 		string
	Value 		string
	Title		string
	IconUrl		string
	Location	string
	Content		string
	BasePrice	string
	Meta	 	string
	Deleted		uint
	CreateTime	string
}

const(
	create_service_primary_table_sql = `CREATE TABLE IF NOT EXISTS service_primary_t(
		uid VARCHAR(64) NOT NULL unique,
		value VARCHAR(64) NOT NULL unique,
		title VARCHAR(64) NULL DEFAULT NULL,
		icon_url VARCHAR(1024) NULL DEFAULT NULL,
		location VARCHAR(64) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		base_price VARCHAR(64) NULL DEFAULT NULL,
		meta VARCHAR(1024) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT 0,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_service_primary = `INSERT INTO service_primary_t (
		uid,
		value,
		title,
		icon_url,
		location,
		content,
		base_price,
		meta,
		deleted) VALUE (?,?,?,?,?,?,?,?,?)`
	query_service_primary = `SELECT * FROM service_primary_t WHERE deleted=0 AND uid=?`
	query_service_primary_by_range = `SELECT * FROM service_primary_t WHERE deleted=0 ORDER BY location LIMIT ? OFFSET ?`
	update_service_primary_by_uid = `UPDATE service_primary_t SET 
		value=?,
		title=?,
		icon_url=?,
		location=?,
		content=?,
		base_price=?,
		meta=? WHERE uid=?`
	update_service_primary_deleted_by_uid = `UPDATE service_primary_t SET deleted=1 WHERE uid=?`
	delete_service_primary_by_uid = `DELETE FROM service_primary_t WHERE uid=?`
)

func CreateServicePrimaryTable() {
	sql := create_service_primary_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create service primary table!");
}

func AddServicePrimay(
	value,
	title,
	iconUrl,
	location,
	content,
	basePrice,
	meta string) error{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_service_primary ,uuid, value, title, iconUrl, location, content, basePrice, meta, 0);
	// Error.CheckErr(err)
	if(nil != err){
		return errors.New("Error: Add Service Primary Error!")
	}
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert service primary success !")
	fmt.Println(aff_nums);
	return nil;
}

func QueryServicePrimary(uid string) *DbServicePrimary{
	
	result := &DbServicePrimary{}
	rows, err := db.Query(query_service_primary, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Value, 
			&result.Title,
			&result.IconUrl,
			&result.Location,
			&result.Content, 
			&result.BasePrice,  
			&result.Meta,
			&result.Deleted,  
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr service primay :")
	fmt.Println(result)
	return result
}

func QueryServicePrimaryByRange(count, offset uint) []DbServicePrimary{

	result := make([]DbServicePrimary, 0)
	rows, err := db.Query(query_service_primary_by_range, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbServicePrimary{}
		err = rows.Scan(
			&item.Uid,
			&item.Value,
			&item.Title,
			&item.IconUrl,
			&item.Location,
			&item.Content,
			&item.BasePrice, 
			&item.Meta,
			&item.Deleted,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func UpdateServicePrimary(
	uid			string,
	value 		string,
	title		string,
	iconUrl		string,
	location	string,
	content		string,
	basePrice	string,
	meta	 	string) {
	//更新数据
	ret, err := db.Exec(update_service_primary_by_uid ,value, title, iconUrl, location, content, basePrice, meta ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update service primary success !")
	fmt.Println(aff_nums);
}

func DeleteServicePrimary(uid string) {
	//更新数据
	ret, err := db.Exec(delete_service_primary_by_uid ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete service primary success !")
	fmt.Println(aff_nums)
}