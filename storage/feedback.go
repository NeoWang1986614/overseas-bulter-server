package storage

import(
	// "database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbFeedback struct{
	Uid 			string
	OrderId 		string
	AuthorId		string
	Content 		string
	IsRead			uint
	Income			float32
	Outgoings		float32
	AccountingDate	string
	UpdateTime		string
	CreateTime		string
}

const(
	create_feedback_table_sql = `CREATE TABLE IF NOT EXISTS feedback_t(
		uid VARCHAR(64) NOT NULL unique,
		order_id VARCHAR(64) NULL DEFAULT NULL,
		author_id VARCHAR(64) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		is_read TINYINT(1) NULL DEFAULT NULL,
		income FLOAT(10,2) NULL DEFAULT NULL,
		outgoings FLOAT(10,2) NULL DEFAULT NULL,
		accounting_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_feedback = `INSERT INTO feedback_t(
		uid,
		order_id,
		author_id,
		content,
		is_read,
		income,
		outgoings,
		accounting_date) value (?,?,?,?,?,?,?,?)`
	query_feedbacks_by_order_id = `SELECT * FROM feedback_t WHERE order_id=? ORDER BY create_time DESC LIMIT ? OFFSET ?`
	update_feedbacks_is_read_by_order_id = `UPDATE feedback_t SET is_read=? WHERE order_id=?`
	query_feedbacks_by_order_id_is_read = `SELECT * FROM feedback_t WHERE order_id=? AND is_read=? LIMIT ? OFFSET ?`
)

func CreateFeedbackTable() {
	sql := create_feedback_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create feedback table!");
	// arr := QueryHousesByRange(5, 0)
}

func AddFeedback(
	orderId 		string,
	authorId 		string,
	content			string,
	income			float32,
	outgoings		float32,
	accountingDate	string) {
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_feedback ,uuid, orderId, authorId, content, 1/*未读*/, income, outgoings, accountingDate);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert feedback success !")
	fmt.Println(aff_nums);
}

func QueryFeedbackByOrderId(count uint, offset uint, orderId string, isFromBackend uint) []DbFeedback{
	
	result := make([]DbFeedback, 0)
	rows, err := db.Query(query_feedbacks_by_order_id, orderId, count ,offset)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		var item = &DbFeedback{}
		err = rows.Scan(
			&item.Uid, 
			&item.OrderId, 
			&item.AuthorId,
			&item.Content, 
			&item.IsRead, 
			&item.Income,
			&item.Outgoings,
			&item.AccountingDate,
			&item.UpdateTime,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}

	if(0==isFromBackend){
		UpdateFeedbacksIsReadByOrderId(orderId);
	}
	
	fmt.Println("queyr feedback :")
	fmt.Println(result)
	return result
}

func UpdateFeedbacksIsReadByOrderId(orderId string){
	//更新数据
	ret, err := db.Exec(update_feedbacks_is_read_by_order_id, 0/*未读*/, orderId);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update feedback to is read success !")
	fmt.Println(aff_nums);
}

func QueryFeedbackByOrderIdIsRead(count uint, offset uint, orderId string, isRead uint) []DbFeedback{
	
	result := make([]DbFeedback, 0)
	rows, err := db.Query(query_feedbacks_by_order_id_is_read, orderId, isRead, count ,offset)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		var item = &DbFeedback{}
		err = rows.Scan(
			&item.Uid, 
			&item.OrderId, 
			&item.AuthorId,
			&item.Content, 
			&item.IsRead, 
			&item.Income,
			&item.Outgoings,
			&item.AccountingDate,
			&item.UpdateTime,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}

	fmt.Println("queyr feedback :")
	fmt.Println(result)
	return result
}