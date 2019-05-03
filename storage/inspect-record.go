package storage

import(
	"fmt"
	// "encoding/json"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbInspectRecord struct{
	Uid 			string
	OrderId			string
	InspectDate 	string
	Inspector 		string	
	Comment 		string
	Config			string
	Area			string
	Deleted			uint
	UpdateTime		string
	CreateTIme		string
}

const(
	create_inspect_record_table_sql = `CREATE TABLE IF NOT EXISTS inspect_record_t(
		uid VARCHAR(64) NOT NULL unique,
		order_id VARCHAR(64) NULL DEFAULT NULL,
		inspect_date VARCHAR(64) NULL DEFAULT NULL,
		inspector VARCHAR(64) NULL DEFAULT NULL,
		comment VARCHAR(64) NULL DEFAULT NULL,
		config VARCHAR(5120) NULL DEFAULT NULL,
		area VARCHAR(5120) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT NULL,
		update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	add_inspect_record = `INSERT INTO inspect_record_t(
		uid,
		order_id,
		inspect_date,
		inspector,
		comment,
		config,
		area,
		deleted) VALUE (?,?,?,?,?,?,?,?)`
	update_inspect_record_by_uid = `UPDATE inspect_record_t SET 
		order_id=?,
		inspect_date=?,
		inspector=?,
		comment=?,
		config=?,
		area=? WHERE uid=?`
	query_inspect_record = `SELECT * FROM inspect_record_t WHERE uid=?`
	query_inspect_record_all_by_order_id = `SELECT COUNT(*) FROM inspect_record_t WHERE order_id=? AND deleted = 0`
	query_inspect_record_by_order_id = `SELECT * FROM inspect_record_t WHERE order_id=? AND deleted = 0 LIMIT ? OFFSET ?`
	update_inspect_record_deleted_by_uid = `UPDATE inspect_record_t SET deleted=1 WHERE uid=?`
)

func CreateInspectRecordTable() {
	sql := create_inspect_record_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create inspect record table");
}

func QueryInspectRecord(uid string) *DbInspectRecord{
	
	ret := &DbInspectRecord{}
	rows, err := db.Query(query_inspect_record, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&ret.Uid,
			&ret.OrderId,
			&ret.InspectDate,
			&ret.Inspector,
			&ret.Comment,
			&ret.Config,
			&ret.Area,
			&ret.Deleted,
			&ret.UpdateTime,
			&ret.CreateTIme)
		Error.CheckErr(err)
	}
	fmt.Println("queyr inspect record :")
	fmt.Println(ret)
	return ret
}

func AddInspectRecord(
	orderId 	string,
	inspectDate string,
	inspector 	string,
	comment 	string,
	config 		string,
	area 		string) string{
	uuid := Uuid.GenerateNextUuid()
	ret, err := db.Exec(add_inspect_record,
		uuid,
		orderId,
		inspectDate,
		inspector,
		comment,
		config,
		area,
		0);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
	return uuid
}

func UpdateInspectRecord(
	uid 		string,
	orderId 	string,
	inspectDate string,
	inspector 	string,
	comment 	string,
	config 		string,
	area 		string) {
	ret, err := db.Exec(update_inspect_record_by_uid,
		orderId,
		inspectDate,
		inspector,
		comment,
		config,
		area,
		uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
}

func DeleteInspectRecordByUid(uid string){
	ret, err := db.Exec(update_inspect_record_deleted_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update deleted inspect record success !")
	fmt.Println(aff_nums);
}

func QueryInspectRecordTotalCountByOrderId(orderId string) uint{
	var count uint = 0
	err := db.QueryRow(query_inspect_record_all_by_order_id, orderId).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryInspectRecordByOrderId(orderId string, offset, count uint) (uint,[]DbInspectRecord){

	total := QueryInspectRecordTotalCountByOrderId(orderId);

	rows, err := db.Query(query_inspect_record_by_order_id, orderId, count, offset)
	defer rows.Close()

	result := make([]DbInspectRecord, 0)
	if err == sql.ErrNoRows {
		return 0, result;
	}
	Error.CheckErr(err)
	for rows.Next() {
		result = append(result, *scanInspectRecordFromRows(rows))
	}
	return total, result;
}

func scanInspectRecordFromRows(rows *sql.Rows) *DbInspectRecord{
	var ret = &DbInspectRecord{}
	err := rows.Scan(
		&ret.Uid,
		&ret.OrderId,
		&ret.InspectDate,
		&ret.Inspector,
		&ret.Comment,
		&ret.Config,
		&ret.Area,
		&ret.Deleted,
		&ret.UpdateTime,
		&ret.CreateTIme)
	Error.CheckErr(err)
	return ret
}