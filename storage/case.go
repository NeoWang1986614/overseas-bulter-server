package storage

import(
	// "database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbCase struct{
	Uid 		string
	Title 		string
	ImageUrl 	string
	Content		string
	Price 		uint
	Level		uint
	CreateTime	string
}

const(
	create_case_table_sql = `CREATE TABLE IF NOT EXISTS case_t(
		uid VARCHAR(64) NOT NULL unique,
		title VARCHAR(64) NULL DEFAULT NULL,
		image_url VARCHAR(2048) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		price INT(64) NULL DEFAULT NULL,
		level TINYINT(1) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_case = `INSERT INTO case_t (uid,title,image_url,content,price,level) value (?,?,?,?,?,?)`
	query_case = `SELECT * FROM case_t WHERE uid=?`
	query_cases_by_level = `SELECT * FROM case_t WHERE level=? LIMIT ? OFFSET ?`
	query_all_cases = `SELECT * FROM case_t`
	update_case_by_uid = `UPDATE case_t SET title=?,image_url=?,content=?,price=?,level=? WHERE uid=?`
	delete_case_by_uid = `DELETE FROM case_t WHERE uid=?`
)

func CreateCaseTable() {
	sql := create_case_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create case table!");
	// arr := QueryHousesByRange(5, 0)
}

func QueryCase(id string) *DbCase{
	
	result := &DbCase{}
	rows, err := db.Query(query_case, id)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid,
			&result.Title,
			&result.ImageUrl,
			&result.Content,
			&result.Price,
			&result.Level,
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr case :")
	fmt.Println(result)
	return result
}

func AddCase(
	title 		string,
	imageUrl 	string,
	content		string,
	price 		uint,
	level		uint) string{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_case ,uuid, title, imageUrl, content, price, level);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert case success !")
	fmt.Println(aff_nums);
	return uuid
}

func UpdateCase(
	uid			string,
	title 		string,
	imageUrl 	string,
	content		string,
	price 		uint,
	level		uint) {
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(update_case_by_uid, title, imageUrl, content, price, level, uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update case success !")
	fmt.Println(aff_nums);
}

func QueryCasesByLevel(count uint, offset uint, level uint) []DbCase{
	
	result := make([]DbCase, 0)
	rows, err := db.Query(query_cases_by_level, level, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbCase{}
		err = rows.Scan(
			&item.Uid,
			&item.Title,
			&item.ImageUrl,
			&item.Content,
			&item.Price,
			&item.Level,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func DeleteCaseByUid(uid string){
	//删除数据
	ret, err := db.Exec(delete_case_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete case success !")
	fmt.Println(aff_nums);
}