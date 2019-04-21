package storage

import(
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbEmployee struct{
	Id string
	UserName string
	Password string
	Nickname string
	AvatarUrl string
	PhoneNumber string
	CreateTime string
}

const(
	create_employee_table_sql = `CREATE TABLE IF NOT EXISTS employee_t(
		id VARCHAR(64) NOT NULL unique,
		user_name VARCHAR(64) NULL DEFAULT NULL,
		nick_name VARCHAR(64) NULL DEFAULT NULL,
		pwd VARCHAR(64) NULL DEFAULT NULL,
		avatar_url VARCHAR(256) NULL DEFAULT NULL,
		phone_number VARCHAR(64) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(id))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	add_employee_if_not_exist = `INSERT INTO employee_t(id, user_name, nick_name, pwd, avatar_url, phone_number) SELECT ?,?,?,?,?,? FROM DUAL WHERE NOT EXISTS (SELECT * FROM employee_t WHERE user_name=?)`
	add_employee = `INSERT INTO employee_t (id, user_name, nick_name, pwd, avatar_url, phone_number) value (?,?,?,?,?,?)`
	query_employee = `SELECT * FROM employee_t WHERE id=?`
	query_employee_by_username = `SELECT * FROM employee_t WHERE user_name=?`
	query_employee_range = `SELECT * FROM employee_t LIMIT ? OFFSET ?`
	query_employee_all = `SELECT * FROM employee_t`
	update_employee = `UPDATE employee_t SET user_name=?,nick_name=?, pwd=?, avatar_url=?, phone_number=? WHERE id=?`
	// update_user_by_uid = `UPDATE user_t SET name=?, phone_number=?,id_card_number=? WHERE uid=?`
	// query_user_by_wx_open_id = `SELECT * FROM user_t WHERE wx_open_id=?`
	// query_user_by_uid = `SELECT * FROM user_t WHERE uid=?`
	// //
	// query_user_by_id_card_number = `SELECT * FROM user_t WHERE id_card_number=?`
	// query_user_by_phone_number = `SELECT * FROM user_t WHERE phone_number=?`
	// query_user_by_name = `SELECT * FROM user_t WHERE name=?`

)

func CreateEmployeeTable() {
	sql := create_employee_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create employee table");
	AddEmployeeIfNotExist("root", "root123", "管理员", "","")
}

func AddEmployeeIfNotExist(
	userName string,
	password string,
	nickname string,
	avatarUrl string,
	phoneNumber string,
	) string{
	uuid := Uuid.GenerateNextUuid()
	ret, err := db.Exec(add_employee_if_not_exist ,uuid, userName, password, nickname, avatarUrl, phoneNumber, userName);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println(aff_nums);
	return uuid
}

func AddEmployee(
	userName string,
	password string,
	nickname string,
	avatarUrl string,
	phoneNumber string,
	) string{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(add_employee ,uuid, userName, password, nickname, avatarUrl, phoneNumber)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected()
	fmt.Println("insert employee success !")
	fmt.Println(aff_nums)
	return uuid
}

func UpdateEmployee(
	id string,
	userName string,
	password string,
	nickname string,
	avatarUrl string,
	phoneNumber string,
	) {
	//更新数据
	ret, err := db.Exec(update_employee , userName, nickname, password, avatarUrl, phoneNumber, id)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update employee success !")
	fmt.Println(aff_nums)
}

func QueryEmployee(id string) *DbEmployee{
	result := &DbEmployee{}
	err := db.QueryRow(query_employee, id).Scan(
		&result.Id,
		&result.UserName,
		&result.Password,
		&result.Nickname,
		&result.AvatarUrl,
		&result.PhoneNumber,
		&result.CreateTime,
	)
	Error.CheckErr(err)
	return result
}

func QueryEmployeeByUserName(userName string) []DbEmployee{
	result := make([]DbEmployee, 0)
	rows, err := db.Query(query_employee_by_username, userName)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		var item = &DbEmployee{}
		err = rows.Scan(
			&item.Id, 
			&item.UserName, 
			&item.Password,
			&item.Nickname, 
			&item.AvatarUrl, 
			&item.PhoneNumber,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	fmt.Println(result)
	return result
}

func QueryEmployeesTotalCount() uint{
	var count uint = 0
	err := db.QueryRow(query_employee_all).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count
}

func QueryEmployees(offset uint, length uint) (uint,[]DbEmployee){

	total := QueryEmployeesTotalCount()
	fmt.Println("total = " , total);

	rows, err := db.Query(query_employee_range, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbEmployee, 0)
	for rows.Next() {
		var item = &DbEmployee{}
		err = rows.Scan(
			&item.Id,
			&item.UserName,
			&item.Password,
			&item.Nickname,
			&item.AvatarUrl,
			&item.PhoneNumber,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total,result;
	
}