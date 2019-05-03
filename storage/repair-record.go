package storage

import(
	"fmt"
	// "encoding/json"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbRepairRecord struct{
	Uid 			string
	OrderId			string
	ReportTime 		string
	RepairTime 		string	
	CompleteTime 	string
	Comment			string
	Cost 			float32
	Status			string
	RelatedImage	string
	Deleted 		uint
	UpdateTime		string
	CreateTIme		string
}

const(
	create_repair_record_table_sql = `CREATE TABLE IF NOT EXISTS repair_record_t(
		uid VARCHAR(64) NOT NULL unique,
		order_id VARCHAR(64) NULL DEFAULT NULL,
		report_time VARCHAR(64) NULL DEFAULT NULL,
		repair_time VARCHAR(64) NULL DEFAULT NULL,
		complete_time VARCHAR(64) NULL DEFAULT NULL,
		comment VARCHAR(5120) NULL DEFAULT NULL,
		cost FLOAT(10,2) NULL DEFAULT NULL,
		status VARCHAR(64) NULL DEFAULT NULL,
		related_image VARCHAR(5120) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT NULL,
		update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	add_repair_record = `INSERT INTO repair_record_t(
		uid,
		order_id,
		report_time,
		repair_time,
		complete_time,
		comment,
		cost,
		status,
		related_image,
		deleted) VALUE (?,?,?,?,?,?,?,?,?,?)`
	update_repair_record_by_uid = `UPDATE repair_record_t SET 
		order_id=?,
		report_time=?,
		repair_time=?,
		complete_time=?,
		comment=?,
		cost=?,
		status=?,
		related_image=? WHERE uid=?`
	query_repair_record = `SELECT * FROM repair_record_t WHERE uid=?`
	query_repair_record_all_by_order_id = `SELECT COUNT(*) FROM repair_record_t WHERE order_id=? AND deleted = 0`
	query_repair_record_by_order_id = `SELECT * FROM repair_record_t WHERE order_id=? AND deleted = 0 LIMIT ? OFFSET ?`
	update_repair_record_deleted_by_uid = `UPDATE repair_record_t SET deleted=1 WHERE uid=?`
)

func CreateRepairRecordTable() {
	sql := create_repair_record_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create repair record table");
}

func QueryRepairRecord(uid string) *DbRepairRecord{
	
	ret := &DbRepairRecord{}
	rows, err := db.Query(query_repair_record, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&ret.Uid,
			&ret.OrderId,
			&ret.ReportTime,
			&ret.RepairTime,
			&ret.CompleteTime,
			&ret.Comment,
			&ret.Cost,
			&ret.Status,
			&ret.RelatedImage,
			&ret.Deleted,
			&ret.UpdateTime,
			&ret.CreateTIme)
		Error.CheckErr(err)
	}
	fmt.Println("queyr repair record :")
	fmt.Println(ret)
	return ret
}

func AddRepairRecord(
	orderId 		string,
	reportTime 		string,
	repairTime 		string,
	completeTime 	string,
	comment 		string,
	cost 			float32,
	status 			string,
	relatedImage 	string) string{
	uuid := Uuid.GenerateNextUuid()
	ret, err := db.Exec(add_repair_record,
		uuid,
		orderId,
		reportTime,
		repairTime,
		completeTime,
		comment,
		cost,
		status,
		relatedImage,
		0);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
	return uuid
}

func UpdateRepairRecord(
	uid 			string,
	orderId 		string,
	reportTime 		string,
	repairTime 		string,
	completeTime 	string,
	comment 		string,
	cost 			float32,
	status 			string,
	relatedImage 	string) {
	ret, err := db.Exec(update_repair_record_by_uid,
		orderId,
		reportTime,
		repairTime,
		completeTime,
		comment,
		cost,
		status,
		relatedImage,
		uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
}

func DeleteRepairRecordByUid(uid string){
	ret, err := db.Exec(update_repair_record_deleted_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update deleted repair record success !")
	fmt.Println(aff_nums);
}

func QueryRepairRecordTotalCountByOrderId(orderId string) uint{
	var count uint = 0
	err := db.QueryRow(query_repair_record_all_by_order_id, orderId).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryRepairRecordByOrderId(orderId string, offset, count uint) (uint,[]DbRepairRecord){

	total := QueryRepairRecordTotalCountByOrderId(orderId);

	rows, err := db.Query(query_repair_record_by_order_id, orderId, count, offset)
	defer rows.Close()

	result := make([]DbRepairRecord, 0)
	if err == sql.ErrNoRows {
		return 0, result;
	}
	Error.CheckErr(err)
	for rows.Next() {
		result = append(result, *scanRepairRecordFromRows(rows))
	}
	return total, result;
}

func scanRepairRecordFromRows(rows *sql.Rows) *DbRepairRecord{
	var ret = &DbRepairRecord{}
	err := rows.Scan(
		&ret.Uid,
		&ret.OrderId,
		&ret.ReportTime,
		&ret.RepairTime,
		&ret.CompleteTime,
		&ret.Comment,
		&ret.Cost,
		&ret.Status,
		&ret.RelatedImage,
		&ret.Deleted,
		&ret.UpdateTime,
		&ret.CreateTIme)
	Error.CheckErr(err)
	return ret
}
