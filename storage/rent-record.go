package storage

import(
	"fmt"
	// "encoding/json"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbRentRecord struct{
	Uid 			string
	OrderId			string
	Income 			string
	Outgoings 		string	
	Balance 		float32
	Comment			string
	TimeRange		string
	AccountingDate	string
	Deleted			uint
	UpdateTime		string
	CreateTIme		string
}

const(
	create_rent_record_table_sql = `CREATE TABLE IF NOT EXISTS rent_record_t(
		uid VARCHAR(64) NOT NULL unique,
		order_id VARCHAR(64) NULL DEFAULT NULL,
		income VARCHAR(5120) NULL DEFAULT NULL,
		outgoings VARCHAR(5120) NULL DEFAULT NULL,
		balance FLOAT(10,2) NULL DEFAULT NULL,
		comment VARCHAR(5120) NULL DEFAULT NULL,
		time_range VARCHAR(64) NULL DEFAULT NULL,
		accounting_date VARCHAR(64) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT NULL,
		update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	add_rent_record = `INSERT INTO rent_record_t(
		uid,
		order_id,
		income,
		outgoings,
		balance,
		comment,
		time_range,
		accounting_date,
		deleted) VALUE (?,?,?,?,?,?,?,?,?)`
	update_rent_record_by_uid = `UPDATE rent_record_t SET 
		order_id=?,
		income=?,
		outgoings=?,
		balance=?,
		comment=?,
		time_range=?,
		accounting_date=? WHERE uid=?`
	query_rent_record = `SELECT * FROM rent_record_t WHERE uid=?`
	query_rent_record_all_by_order_id = `SELECT COUNT(*) FROM rent_record_t WHERE order_id=? AND deleted = 0`
	query_rent_record_by_order_id = `SELECT * FROM rent_record_t WHERE order_id=? AND deleted = 0 LIMIT ? OFFSET ?`
	update_rent_record_deleted_by_uid = `UPDATE rent_record_t SET deleted=1 WHERE uid=?`
)

func CreateRentRecordTable() {
	sql := create_rent_record_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create rent record table");
}

func QueryRentRecord(uid string) *DbRentRecord{
	
	ret := &DbRentRecord{}
	rows, err := db.Query(query_rent_record, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&ret.Uid,
			&ret.OrderId,
			&ret.Income,
			&ret.Outgoings,
			&ret.Balance,
			&ret.Comment,
			&ret.TimeRange,
			&ret.AccountingDate,
			&ret.Deleted,
			&ret.UpdateTime,
			&ret.CreateTIme)
		Error.CheckErr(err)
	}
	fmt.Println("queyr rent record :")
	fmt.Println(ret)
	return ret
}

func AddRentRecord(
	orderId string,
	income string,
	outgoings string,
	balance float32,
	comment string,
	timeRange string,
	accountingDate string) string{
	uuid := Uuid.GenerateNextUuid()
	ret, err := db.Exec(add_rent_record,
		uuid,
		orderId,
		income,
		outgoings,
		balance,
		comment,
		timeRange,
		accountingDate,
		0);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
	return uuid
}

func UpdateRentRecord(
	uid string,
	orderId string,
	income string,
	outgoings string,
	balance float32,
	comment string,
	timeRange string,
	accountingDate string) {
	ret, err := db.Exec(update_rent_record_by_uid,
		orderId,
		income,
		outgoings,
		balance,
		comment,
		timeRange,
		accountingDate,
		uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
}

func DeleteRentRecordByUid(uid string){
	ret, err := db.Exec(update_rent_record_deleted_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update deleted rent record success !")
	fmt.Println(aff_nums);
}

func QueryRentRecordTotalCountByOrderId(orderId string) uint{
	var count uint = 0
	err := db.QueryRow(query_rent_record_all_by_order_id, orderId).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryRentRecordByOrderId(orderId string, offset, count uint) (uint,[]DbRentRecord){

	total := QueryRentRecordTotalCountByOrderId(orderId);

	rows, err := db.Query(query_rent_record_by_order_id, orderId, count, offset)
	defer rows.Close()

	result := make([]DbRentRecord, 0)
	if err == sql.ErrNoRows {
		return 0, result;
	}
	Error.CheckErr(err)
	for rows.Next() {
		result = append(result, *scanRentRecordFromRows(rows))
	}
	return total, result;
}

func scanRentRecordFromRows(rows *sql.Rows) *DbRentRecord{
	var ret = &DbRentRecord{}
	err := rows.Scan(
		&ret.Uid,
		&ret.OrderId,
		&ret.Income,
		&ret.Outgoings,
		&ret.Balance,
		&ret.Comment,
		&ret.TimeRange,
		&ret.AccountingDate,
		&ret.Deleted,
		&ret.UpdateTime,
		&ret.CreateTIme)
	Error.CheckErr(err)
	return ret
}