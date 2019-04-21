package storage

import(
	// "database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbUser struct{
	Uid string
	Name string
	WxOpenId string
	WxSessionKey string
	PhoneNumber string
	IdCardNumber string
	CreateTime string
}

const(
	create_user_table_sql = `CREATE TABLE IF NOT EXISTS user_t(
		uid VARCHAR(64) NOT NULL unique,
		name VARCHAR(64) NULL DEFAULT NULL,
		wx_open_id VARCHAR(64) NULL DEFAULT NULL unique,
		wx_session_key VARCHAR(64) NULL DEFAULT NULL,
		phone_number VARCHAR(64) NULL DEFAULT NULL,
		id_card_number VARCHAR(64) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	update_user_by_wx_open_id = `INSERT INTO user_t (uid,name, wx_open_id,wx_session_key,phone_number, id_card_number) value (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE wx_session_key=?`
	update_user_by_uid = `UPDATE user_t SET name=?, phone_number=?,id_card_number=? WHERE uid=?`
	query_user_by_wx_open_id = `SELECT * FROM user_t WHERE wx_open_id=?`
	query_user_by_uid = `SELECT * FROM user_t WHERE uid=?`
	//
	query_user_by_id_card_number = `SELECT * FROM user_t WHERE id_card_number=?`
	query_user_by_phone_number = `SELECT * FROM user_t WHERE phone_number=?`
	query_user_by_name = `SELECT * FROM user_t WHERE name=?`

)

func CreateUserTable() {
	sql := create_user_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create user table");
}

func UpdateUserByWxOpenId(wxOpenId string, wxSessionKey string) {
	uuid := Uuid.GenerateNextUuid()
	//更新数据
	ret, err := db.Exec(update_user_by_wx_open_id, uuid, "", wxOpenId, wxSessionKey,"","", wxSessionKey);
	Error.CheckErr(err)
    aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
	QueryUserByWxOpenId(wxOpenId)
}

func UpdateUserByUid(uid string, name string, phoneNumber string, idCardNumber string) {
	//更新数据
	ret, err := db.Exec(update_user_by_uid, name, phoneNumber, idCardNumber, uid);
	Error.CheckErr(err)
    aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
}

func QueryUserByWxOpenId(wxOpenId string) *DbUser{
	
	result := &DbUser{}
	rows, err := db.Query(query_user_by_wx_open_id, wxOpenId)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Name, 
			&result.WxOpenId, 
			&result.WxSessionKey,  
			&result.PhoneNumber, 
			&result.IdCardNumber, 
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr user :")
	fmt.Println(result)
	return result
}

func QueryUserByUid(uid string) *DbUser{
	fmt.Println("QueryUserByUid , uid = ", uid)
	result := &DbUser{}
	rows, err := db.Query(query_user_by_uid, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Name, 
			&result.WxOpenId, 
			&result.WxSessionKey,  
			&result.PhoneNumber, 
			&result.IdCardNumber, 
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr user :")
	fmt.Println(result)
	return result
}

func QueryUserByIdCardNumber(idCardNumber string) []DbUser{
	result := make([]DbUser, 0)
	rows, err := db.Query(query_user_by_id_card_number, idCardNumber)
	defer rows.Close()
	Error.CheckErr(err)
	
	for rows.Next() {
		var item = &DbUser{}
		err = rows.Scan(
			&item.Uid, 
			&item.Name, 
			&item.WxOpenId, 
			&item.WxSessionKey,  
			&item.PhoneNumber, 
			&item.IdCardNumber, 
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func QueryUserByPhoneNumber(phoneNumber string) []DbUser{
	result := make([]DbUser, 0)
	rows, err := db.Query(query_user_by_phone_number, phoneNumber)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		var item = &DbUser{}
		err = rows.Scan(
			&item.Uid, 
			&item.Name, 
			&item.WxOpenId, 
			&item.WxSessionKey,  
			&item.PhoneNumber, 
			&item.IdCardNumber, 
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func QueryUserByName(name string) []DbUser{
	result := make([]DbUser, 0)
	rows, err := db.Query(query_user_by_name, name)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		var item = &DbUser{}
		err = rows.Scan(
			&item.Uid, 
			&item.Name, 
			&item.WxOpenId, 
			&item.WxSessionKey,  
			&item.PhoneNumber, 
			&item.IdCardNumber, 
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}