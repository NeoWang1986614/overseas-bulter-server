package storage

import(
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbWechat struct{
	Id 			string
	AccessToken string
	ExpiresIn 	uint
	UpdateTime 	string
}

const(
	create_wechat_table_sql = `CREATE TABLE IF NOT EXISTS wechat_t(
		id VARCHAR(64) NOT NULL unique,
		access_token VARCHAR(1024) NULL DEFAULT NULL,
		expires_in INT(64) NULL DEFAULT NULL,
		update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY(id))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	add_wechat = `INSERT INTO wechat_t(id, access_token, expires_in) VALUE (?,?,?)`
	query_wechat_all = `SELECT * FROM wechat_t`
	update_wechat_by_id = `UPDATE wechat_t SET access_token=?, expires_in=? WHERE id=?`
)

func CreateWechatTable() {
	sql := create_wechat_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create wechat table");
}

func AddWechat(
	accessToken string,
	expiresIn uint) string{
	uuid := Uuid.GenerateNextUuid()
	ret, err := db.Exec(add_wechat ,uuid, accessToken, expiresIn);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
	return uuid
}

func UpdateWechat(
	id string,
	accessToken string,
	expiresIn uint) {
	//更新数据
	ret, err := db.Exec(update_wechat_by_id, accessToken, expiresIn, id)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update wechat success !")
	fmt.Println(aff_nums)
}

func QueryWechatAll() []DbWechat{

	rows, err := db.Query(query_wechat_all)
	defer rows.Close()
	if err == sql.ErrNoRows {
		return []DbWechat{}
	}else{
		Error.CheckErr(err)
	}
	result := make([]DbWechat, 0)
	for rows.Next() {
		result = append(result, *scanWechatItemFromRows(rows))
	}
	// fmt.Println(result)
	return result;
}

func scanWechatItemFromRows(rows *sql.Rows) *DbWechat{
	var ret = &DbWechat{}
	err := rows.Scan(
		&ret.Id,
		&ret.AccessToken,
		&ret.ExpiresIn,
		&ret.UpdateTime)
	Error.CheckErr(err)
	return ret;
}