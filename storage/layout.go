package storage

import(
	"fmt"
	"errors"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbLayout struct{
	Uid 		string
	Value 		string
	Title		string
	Location	string
	Content		string
	Meta	 	string
	Deleted		uint
	CreateTime	string
}

const(
	create_layout_table_sql = `CREATE TABLE IF NOT EXISTS layout_t(
		uid VARCHAR(64) NOT NULL unique,
		value VARCHAR(64) NOT NULL unique,
		title VARCHAR(64) NULL DEFAULT NULL,
		location VARCHAR(64) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		meta VARCHAR(1024) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT 0,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_layout = `INSERT INTO layout_t (
		uid,
		value,
		title,
		location,
		content,
		meta,
		deleted) VALUE (?,?,?,?,?,?,?)`
	query_layout = `SELECT * FROM layout_t WHERE deleted=0 AND uid=?`
	query_layout_by_range = `SELECT * FROM layout_t WHERE deleted=0 ORDER BY location LIMIT ? OFFSET ?`
	update_layout_by_uid = `UPDATE layout_t SET 
		value=?,
		title=?,
		location=?,
		content=?,
		meta=? WHERE uid=?`
	update_layout_deleted_by_uid = `UPDATE layout_t SET deleted=1 WHERE uid=?`
	delete_layout_deleted_by_uid = `DELETE FROM layout_t WHERE uid=?`
)

func CreateLayoutTable() {
	sql := create_layout_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create layout table!");
}

func AddLayout(
	value,
	title,
	location,
	content,
	meta string) error{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_layout ,uuid, value, title, location, content, meta, 0);
	// Error.CheckErr(err)
	if(nil != err){
		return errors.New("Error: Add Layout Error!")
	}
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert layout success !")
	fmt.Println(aff_nums);
	return nil;
}

func QueryLayout(uid string) *DbLayout{
	
	result := &DbLayout{}
	rows, err := db.Query(query_layout, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Value, 
			&result.Title,
			&result.Location,
			&result.Content,  
			&result.Meta,
			&result.Deleted,  
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr layout :")
	fmt.Println(result)
	return result
}

func QueryLayoutByRange(count, offset uint) []DbLayout{

	result := make([]DbLayout, 0)
	rows, err := db.Query(query_layout_by_range, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbLayout{}
		err = rows.Scan(
			&item.Uid,
			&item.Value,
			&item.Title,
			&item.Location,
			&item.Content,
			&item.Meta,
			&item.Deleted,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func UpdateLayout(
	uid			string,
	value 		string,
	title		string,
	location	string,
	content		string,
	meta	 	string) {
	//更新数据
	ret, err := db.Exec(update_layout_by_uid ,value, title, location, content, meta ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update layout success !")
	fmt.Println(aff_nums);
}

func DeleteLayout(uid string) {
	//更新数据
	ret, err := db.Exec(delete_layout_deleted_by_uid ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete layout success !")
	fmt.Println(aff_nums);
}