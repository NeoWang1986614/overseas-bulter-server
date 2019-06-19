package storage

import(
	"fmt"
	// "encoding/json"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbBillingRecord struct{
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
	create_billing_record_table_sql = `CREATE TABLE IF NOT EXISTS billing_record_t(
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
	add_billing_record = `INSERT INTO billing_record_t(
		uid,
		order_id,
		income,
		outgoings,
		balance,
		comment,
		time_range,
		accounting_date,
		deleted) VALUE (?,?,?,?,?,?,?,?,?)`
	update_billing_record_by_uid = `UPDATE billing_record_t SET 
		order_id=?,
		income=?,
		outgoings=?,
		balance=?,
		comment=?,
		time_range=?,
		accounting_date=? WHERE uid=?`
	query_billing_record = `SELECT * FROM billing_record_t WHERE uid=?`
	query_billing_record_all_by_order_id = `SELECT COUNT(*) FROM billing_record_t WHERE order_id=? AND deleted = 0`
	query_billing_record_by_order_id = `SELECT * FROM billing_record_t WHERE order_id=? AND deleted = 0 LIMIT ? OFFSET ?`
	update_billing_record_deleted_by_uid = `UPDATE billing_record_t SET deleted=1 WHERE uid=?`
)

func CreateBillingRecordTable() {
	sql := create_billing_record_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create billing record table");
}

func QueryBillingRecord(uid string) *DbBillingRecord{
	
	ret := &DbBillingRecord{}
	rows, err := db.Query(query_billing_record, uid)
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
	fmt.Println("queyr billing record :")
	fmt.Println(ret)
	return ret
}

func AddBillingRecord(
	orderId string,
	income string,
	outgoings string,
	balance float32,
	comment string,
	timeRange string,
	accountingDate string) string{
	uuid := Uuid.GenerateNextUuid()
	ret, err := db.Exec(add_billing_record,
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

func UpdateBillingRecord(
	uid string,
	orderId string,
	income string,
	outgoings string,
	balance float32,
	comment string,
	timeRange string,
	accountingDate string) {
	ret, err := db.Exec(update_billing_record_by_uid,
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

func DeleteBillingRecordByUid(uid string){
	ret, err := db.Exec(update_billing_record_deleted_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update deleted billing record success !")
	fmt.Println(aff_nums);
}

func QueryBillingRecordTotalCountByOrderId(orderId string) uint{
	var count uint = 0
	err := db.QueryRow(query_billing_record_all_by_order_id, orderId).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryBillingRecordByOrderId(orderId string, offset, count uint) (uint,[]DbBillingRecord){

	total := QueryBillingRecordTotalCountByOrderId(orderId);

	rows, err := db.Query(query_billing_record_by_order_id, orderId, count, offset)
	defer rows.Close()

	result := make([]DbBillingRecord, 0)
	if err == sql.ErrNoRows {
		return 0, result;
	}
	Error.CheckErr(err)
	for rows.Next() {
		result = append(result, *scanBillingRecordFromRows(rows))
	}
	return total, result;
}

func scanBillingRecordFromRows(rows *sql.Rows) *DbBillingRecord{
	var ret = &DbBillingRecord{}
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