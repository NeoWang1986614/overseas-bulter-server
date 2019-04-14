// Uid 				string	`json:"uid"`
// 	Type				string	`json:"type"`
// 	Content 			string  `json:"title"` 
// 	HouseId 			string	`json:"house_id"`
// 	Price				uint	`json:"price"`
// 	Status				string  `json:"status"`
// 	PlacerId			string	`json:"placer_id"`
// 	AccepterId			string	`json:"accepter_id"`
// 	Time				string	`json:"time"`

package storage

import(
	// "database/sql"
	"fmt"
	"strings"
	// _"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbOrder struct{//13
	Uid 			string
	Type 			string
	Content 		string
	HouseCountry	string
	HouseProvince	string
	HouseCity		string
	HouseAddress	string
	HouseLayout		string
	Price 			uint
	Status			string
	PlacerId		string
	AccepterId		string
	CreateTime		string
}

const(
	create_order_table_sql = `CREATE TABLE IF NOT EXISTS order_t(
		uid VARCHAR(64) NOT NULL unique,
		type VARCHAR(64) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		house_country VARCHAR(64) NULL DEFAULT NULL,
		house_province VARCHAR(64) NULL DEFAULT NULL,
		house_city VARCHAR(64) NULL DEFAULT NULL,
		house_address VARCHAR(1024) NULL DEFAULT NULL,
		house_layout VARCHAR(64) NULL DEFAULT NULL,
		price INT(64) NULL DEFAULT NULL,
		status VARCHAR(64) NULL DEFAULT NULL,
		placer_id VARCHAR(64) NULL DEFAULT NULL,
		accepter_id VARCHAR(64) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_order = `INSERT INTO order_t (uid,type,content,house_country,house_province, house_city, house_address, house_layout,price,status,placer_id, accepter_id) value (?,?,?,?,?,?,?,?,?,?,?,?)`
	query_orders_by_status_placer = `SELECT * FROM order_t WHERE status=? AND placer_id=? LIMIT ? OFFSET ?`
	query_order = `SELECT * FROM order_t WHERE uid=?`
	update_order_by_id = `UPDATE order_t SET type=?,content=?,house_country=?,house_province=?, house_city=?, house_address=?, house_layout=?,price=?,status=?,placer_id=?,accepter_id=? WHERE uid=?`
	delete_order_by_id = `DELETE FROM order_t WHERE uid=?`

	query_order_by_user_ids = `SELECT * FROM order_t WHERE placer_id IN (%s) LIMIT ? OFFSET ?`
	query_order_count_by_user_ids = `SELECT COUNT(*) FROM order_t WHERE placer_id IN (%s)`
	query_order_before_time = `SELECT * FROM order_t WHERE create_time <= ? LIMIT ? OFFSET ?`
	query_order_count_before_time = `SELECT COUNT(*) FROM order_t WHERE create_time <= ?`
	query_order_after_time = `SELECT * FROM order_t WHERE create_time >= ? LIMIT ? OFFSET ?`
	query_order_count_after_time = `SELECT COUNT(*) FROM order_t WHERE create_time >= ?`
	query_order_range_time = `SELECT * FROM order_t WHERE create_time >= ? AND create_time <= ? LIMIT ? OFFSET ?`
	query_order_count_range_time = `SELECT COUNT(*) FROM order_t WHERE create_time >= ? AND create_time <= ?`
	query_order_by_status_group = `SELECT * FROM order_t WHERE status IN (%s) LIMIT ? OFFSET ?`
	query_order_count_by_status_group = `SELECT COUNT(*) FROM order_t WHERE status IN (%s)`
	query_order_by_address = `SELECT * FROM order_t WHERE house_country=%s AND house_province=%s AND house_city=%s LIMIT ? OFFSET ?`
	query_order_count_by_address = `SELECT COUNT(*) FROM order_t WHERE house_country=%s AND house_province=%s AND house_city=%s`
	query_order_by_layout_group = `SELECT * FROM order_t WHERE house_layout IN (%s) LIMIT ? OFFSET ?`
	query_order_count_by_layout_group = `SELECT COUNT(*) FROM order_t WHERE house_layout IN (%s)`

	query_order_below_price = `SELECT * FROM order_t WHERE price <= ? LIMIT ? OFFSET ?`
	query_order_count_below_price = `SELECT COUNT(*) FROM order_t WHERE price <= ?`

	query_order_above_price = `SELECT * FROM order_t WHERE price >= ? LIMIT ? OFFSET ?`
	query_order_count_above_price = `SELECT COUNT(*) FROM order_t WHERE price >= ?`

	query_order_range_price = `SELECT * FROM order_t WHERE price >= ? AND price <= ? LIMIT ? OFFSET ?`
	query_order_count_range_price = `SELECT COUNT(*) FROM order_t WHERE price >= ? AND price <= ?`
)

func CreateOrderTable() {
	sql := create_order_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create order table!");
	// arr := QueryHousesByRange(5, 0)
}

func QueryOrder(id string) *DbOrder{
	
	result := &DbOrder{}
	rows, err := db.Query(query_order, id)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Type, 
			&result.Content,  
			&result.HouseCountry, 
			&result.HouseProvince, 
			&result.HouseCity, 
			&result.HouseAddress, 
			&result.HouseLayout,  
			&result.Price, 
			&result.Status, 
			&result.PlacerId, 
			&result.AccepterId, 
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr order :")
	fmt.Println(result)
	return result
}

func AddOrder(
	orderType 		string,
	content 		string,
	houseCountry	string,
	houseProvince	string,
	houseCity		string,
	houseAddress	string,
	houseLayout		string,
	price 			uint,
	status			string,
	placerId		string,
	accepterId		string) string{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_order ,uuid, orderType, content, houseCountry, houseProvince, houseCity, houseAddress, houseLayout, price, status, placerId, accepterId);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert order success !")
	fmt.Println(aff_nums);
	return uuid;
}

func UpdateOrder(
	uid 			string,
	orderType 		string,
	content 		string,
	houseCountry	string,
	houseProvince	string,
	houseCity		string,
	houseAddress	string,
	houseLayout		string,
	price 			uint,
	status			string,
	placerId		string,
	accepterId		string){
	
		fmt.Println("pay status ", status);
	ret, err := db.Exec(update_order_by_id, 
		orderType, 
		content, 
		houseCountry,
		houseProvince,
		houseCity,
		houseAddress,
		houseLayout,
		price, 
		status, 
		placerId, 
		accepterId,
		uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update order success !")
	fmt.Println(aff_nums)
}

func QueryOrdersByStatusPlacerId(count uint, offset uint, status string, placerId string) []DbOrder{
	
	result := make([]DbOrder, 0)
	rows, err := db.Query(query_orders_by_status_placer, status, placerId, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func DeleteOrderByUid(uid string){
	//删除数据
	ret, err := db.Exec(delete_order_by_id ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete order success !")
	fmt.Println(aff_nums);
}

func QueryOrderTotalCountByUsers(userIds []string) uint{
	var temp = strings.Join(userIds, ",")
	querySql := fmt.Sprintf(query_order_count_by_user_ids, temp)
	fmt.Println(querySql)

	var count uint = 0
	err := db.QueryRow(querySql).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderByUsers(userIds []string, offset uint, length uint) (uint,[]DbOrder){

	total := QueryOrderTotalCountByUsers(userIds);
	fmt.Println("total = " , total);

	var temp = strings.Join(userIds, ",")
	querySql := fmt.Sprintf(query_order_by_user_ids, temp)
	fmt.Println(querySql)

	rows, err := db.Query(querySql, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total,result;
	
}

func QueryOrderTotalCountBeforeTime(time string,) uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_before_time, time).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderBeforeTime(time string, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountBeforeTime(time);
	fmt.Println("total = " , total);

	rows, err := db.Query(query_order_before_time, time, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total,result;
	
}

func QueryOrderTotalCountAfterTime(time string,) uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_after_time, time).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderAfterTime(time string, offset uint, length uint) (uint,[]DbOrder){

	total := QueryOrderTotalCountAfterTime(time);
	fmt.Println("total = " , total);

	rows, err := db.Query(query_order_after_time, time, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountRangeTime(fromTime string, toTime string) uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_range_time, fromTime, toTime).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderRangeTime(fromTime string, toTime string, offset uint, length uint) (uint,[]DbOrder){

	total := QueryOrderTotalCountRangeTime(fromTime, toTime);
	fmt.Println("total = " , total);

	rows, err := db.Query(query_order_range_time, fromTime, toTime, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountByStatusGroup(statuses []string,) uint{

	var temp = "'" + strings.Join(statuses, "','") + "'"
	querySql := fmt.Sprintf(query_order_count_by_status_group, temp)
	fmt.Println(querySql)

	var count uint = 0
	err := db.QueryRow(querySql).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderByStatusGroup(statuses []string, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountByStatusGroup(statuses);
	fmt.Println("total = " , total);

	var temp = "'" + strings.Join(statuses, "','") + "'"
	querySql := fmt.Sprintf(query_order_by_status_group, temp)
	fmt.Println(querySql)

	rows, err := db.Query(querySql, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountByAddress(countrySubSql string, provinceSubSql string, citySubSql string,) uint{

	querySql := fmt.Sprintf(query_order_count_by_address, countrySubSql, provinceSubSql, citySubSql)
	fmt.Println(querySql)

	var count uint = 0
	err := db.QueryRow(querySql).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderByAddress(country string, province string, city string, offset uint, length uint) (uint, []DbOrder){

	countrySubSql := "'" + country + "'"
	if(0 == len(country)) {
		countrySubSql = "ANY(SELECT house_country FROM order_t)"
	}

	provinceSubSql := "'" + province + "'"
	if(0 == len(province)) {
		provinceSubSql = "ANY(SELECT house_province FROM order_t)"
	}

	citySubSql := "'" + city + "'"
	if(0 == len(city)) {
		citySubSql = "ANY(SELECT house_city FROM order_t)"
	}

	total := QueryOrderTotalCountByAddress(countrySubSql, provinceSubSql, citySubSql);
	fmt.Println("total = " , total);

	fullSql := fmt.Sprintf(query_order_by_address, countrySubSql, provinceSubSql, citySubSql);

	fmt.Println(fullSql)

	rows, err := db.Query(fullSql,length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountByLayoutGroup(layouts []string) uint{

	var temp = "'" + strings.Join(layouts, "','") + "'"
	querySql := fmt.Sprintf(query_order_count_by_layout_group, temp)
	fmt.Println(querySql)
	var count uint = 0
	err := db.QueryRow(querySql).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderByLayoutGroup(layouts []string, offset uint, length uint) (uint,[]DbOrder){

	total := QueryOrderTotalCountByLayoutGroup(layouts)
	fmt.Println("total = " , total);

	var temp = "'" + strings.Join(layouts, "','") + "'"
	querySql := fmt.Sprintf(query_order_by_layout_group, temp)
	fmt.Println(querySql)

	rows, err := db.Query(querySql, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountBelowPrice(price uint) uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_below_price, price).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderBelowPrice(price uint, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountBelowPrice(price)
	fmt.Println("total = " , total);

	rows, err := db.Query(query_order_below_price, price, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountAbovePrice(price uint) uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_above_price, price).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderAbovePrice(price uint, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountAbovePrice(price)
	fmt.Println("total = " , total);

	rows, err := db.Query(query_order_above_price, price, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountRangePrice(fromPrice uint, toPrice uint) uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_range_price, fromPrice, toPrice).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderRangePrice(fromPrice uint, toPrice uint, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountRangePrice(fromPrice, toPrice)
	fmt.Println("total = " , total);

	rows, err := db.Query(query_order_range_price, fromPrice, toPrice, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		var item = &DbOrder{}
		err = rows.Scan(
			&item.Uid,
			&item.Type,
			&item.Content,
			&item.HouseCountry,
			&item.HouseProvince,
			&item.HouseCity,
			&item.HouseAddress,
			&item.HouseLayout,
			&item.Price,
			&item.Status,
			&item.PlacerId,
			&item.AccepterId,
			&item.CreateTime)

		Error.CheckErr(err)
		result = append(result, *item)
	}
	// fmt.Println(result)
	return total, result;
	
}
