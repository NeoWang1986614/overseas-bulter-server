package storage

import(
	"database/sql"
	"fmt"
	"strings"
	// _"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbOrder struct{//13
	Uid 				string
	OrderType 			string
	Content 			string
	HouseId				string
	HouseNation			string
	HouseAdLevel1		string
	HouseAdLevel2		string
	HouseAdLevel3		string
	HouseStreetName		string
	HouseStreetNum		string
	HouseBuildingNum	string
	HouseRoomNum		string
	HouseLayout			string
	HouseArea			float32
	Price 				uint
	Status				string
	PlacerId			string
	AccepterId			string
	WxPrepayId			string
	Meta				string
	Deleted				uint
	CreateTime			string
}

const(
	create_order_table_sql = `CREATE TABLE IF NOT EXISTS order_t(
		uid VARCHAR(64) NOT NULL unique,
		order_type VARCHAR(64) NULL DEFAULT NULL,
		content VARCHAR(2048) NULL DEFAULT NULL,
		house_id VARCHAR(64) NULL DEFAULT NULL,
		house_nation VARCHAR(64) NULL DEFAULT NULL,
		house_ad_level_1 VARCHAR(64) NULL DEFAULT NULL,
		house_ad_level_2 VARCHAR(64) NULL DEFAULT NULL,
		house_ad_level_3 VARCHAR(64) NULL DEFAULT NULL,
		house_street_name VARCHAR(64) NULL DEFAULT NULL,
		house_street_num VARCHAR(64) NULL DEFAULT NULL,
		house_building_num VARCHAR(64) NULL DEFAULT NULL,
		house_room_num VARCHAR(64) NULL DEFAULT NULL,
		house_layout VARCHAR(64) NULL DEFAULT NULL,
		house_area FLOAT(10,2) NULL DEFAULT NULL,
		price INT(64) NULL DEFAULT NULL,
		status VARCHAR(64) NULL DEFAULT NULL,
		placer_id VARCHAR(64) NULL DEFAULT NULL,
		accepter_id VARCHAR(64) NULL DEFAULT NULL,
		wx_prepay_id VARCHAR(64) NULL DEFAULT NULL,
		meta VARCHAR(1024) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT 0,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_order = `INSERT INTO order_t (
		uid,
		order_type,
		content,
		house_id,
		house_nation,
		house_ad_level_1,
		house_ad_level_2,
		house_ad_level_3,
		house_street_name,
		house_street_num,
		house_building_num,
		house_room_num,
		house_layout,
		house_area,
		price,
		status,
		placer_id,
		accepter_id,
		wx_prepay_id,
		meta,
		deleted) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	query_orders_by_status_placer = `SELECT * FROM order_t WHERE status=? AND placer_id=? AND deleted=0 LIMIT ? OFFSET ?`
	query_order = `SELECT * FROM order_t WHERE uid=? AND deleted=0`
	update_order_by_id = `UPDATE order_t SET 
		order_type=?,
		content=?,
		house_id=?,
		house_nation=?,
		house_ad_level_1=?,
		house_ad_level_2=?,
		house_ad_level_3=?,
		house_street_name=?,
		house_street_num=?,
		house_building_num=?,
		house_room_num=?,
		house_layout=?,
		house_area=?,
		price=?,
		status=?,
		placer_id=?,
		accepter_id=?,
		meta=? WHERE uid=?`
	update_order_prepay_id_by_id = `UPDATE order_t temp SET temp.wx_prepay_id='%s' WHERE uid='%s'`
	update_order_status_by_id = `UPDATE order_t SET status=? WHERE uid=?`
	// delete_order_by_id = `DELETE FROM order_t WHERE uid=?`
	update_order_deleted_by_id = `UPDATE order_t SET deleted=1 WHERE uid=?`

	

	query_order_by_user_ids = `SELECT * FROM order_t WHERE placer_id IN (%s) AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_by_user_ids = `SELECT COUNT(*) FROM order_t WHERE placer_id IN (%s) AND deleted=0`
	query_order_before_time = `SELECT * FROM order_t WHERE create_time <= ? AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_before_time = `SELECT COUNT(*) FROM order_t WHERE create_time <= ? AND deleted=0`
	query_order_after_time = `SELECT * FROM order_t WHERE create_time >= ? AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_after_time = `SELECT COUNT(*) FROM order_t WHERE create_time >= ? AND deleted=0`
	query_order_range_time = `SELECT * FROM order_t WHERE create_time >= ? AND create_time <= ? AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_range_time = `SELECT COUNT(*) FROM order_t WHERE create_time >= ? AND create_time <= ? AND deleted=0`
	query_order_by_status_group = `SELECT * FROM order_t WHERE status IN (%s) AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_by_status_group = `SELECT COUNT(*) FROM order_t WHERE status IN (%s)  AND deleted=0`
	query_order_by_address = `SELECT * FROM order_t WHERE house_nation=%s AND house_ad_level_1=%s AND house_ad_level_2=%s  AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_by_address = `SELECT COUNT(*) FROM order_t WHERE house_nation=%s AND house_ad_level_1=%s AND house_ad_level_2=%s  AND deleted=0 `
	query_order_by_layout_group = `SELECT * FROM order_t WHERE house_layout IN (%s)  AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_by_layout_group = `SELECT COUNT(*) FROM order_t WHERE house_layout IN (%s) AND deleted=0`
	query_order_below_price = `SELECT * FROM order_t WHERE price <= ?  AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_below_price = `SELECT COUNT(*) FROM order_t WHERE price <= ?  AND deleted=0 `
	query_order_above_price = `SELECT * FROM order_t WHERE price >= ?  AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_above_price = `SELECT COUNT(*) FROM order_t WHERE price >= ?  AND deleted=0 `
	query_order_range_price = `SELECT * FROM order_t WHERE price >= ? AND price <= ?  AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_range_price = `SELECT COUNT(*) FROM order_t WHERE price >= ? AND price <= ?  AND deleted=0 `
	
	query_order_by_order_type_group = `SELECT * FROM order_t WHERE order_type IN (%s)  AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_by_order_type_group = `SELECT COUNT(*) FROM order_t WHERE order_type IN (%s)  AND deleted=0 `

	query_orders = `SELECT * FROM order_t WHERE deleted=0 LIMIT ? OFFSET ?`
	query_order_count_all = `SELECT COUNT(*) FROM order_t WHERE deleted=0`

	query_order_by_order_type_group_order_status = `SELECT * FROM order_t WHERE order_type IN (%s) AND status=? AND deleted=0 LIMIT ? OFFSET ?`
	query_order_count_by_order_type_group_order_status = `SELECT COUNT(*) FROM order_t WHERE order_type IN (%s) AND status=? AND deleted=0 `

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
			&result.OrderType, 
			&result.Content,
			&result.HouseId,
			&result.HouseNation, 
			&result.HouseAdLevel1, 
			&result.HouseAdLevel2,
			&result.HouseAdLevel3,
			&result.HouseStreetName,
			&result.HouseStreetNum,
			&result.HouseBuildingNum,
			&result.HouseRoomNum,  
			&result.HouseLayout,
			&result.HouseArea,
			&result.Price, 
			&result.Status, 
			&result.PlacerId, 
			&result.AccepterId,
			&result.WxPrepayId,
			&result.Meta,
			&result.Deleted,
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr order :")
	fmt.Println(result)
	return result
}

func AddOrder(
	orderType 			string,
	content 			string,
	houseId				string,
	houseNation			string,
	houseAdLevel1		string,
	houseAdLevel2		string,
	houseAdLevel3		string,
	houseStreetName		string,
	houseStreetNum		string,
	houseBuildingNum 	string,
	HouseRoomNum	 	string,
	houseLayout			string,
	houseArea			float32,
	price 				uint,
	status				string,
	placerId			string,
	accepterId			string,
	wxPrepayId			string,
	meta				string) string{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_order ,
		uuid,
		orderType,
		content,
		houseId,
		houseNation,
		houseAdLevel1,
		houseAdLevel2,
		houseAdLevel3,
		houseStreetName,
		houseStreetNum,
		houseBuildingNum,
		HouseRoomNum,
		houseLayout,
		houseArea,
		price,
		status,
		placerId,
		accepterId,
		wxPrepayId,
		meta,
		0);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert order success !")
	fmt.Println(aff_nums);
	return uuid;
}

func UpdateOrder(
	uid 				string,
	orderType 			string,
	content 			string,
	houseId 			string,
	houseNation			string,
	houseAdLevel1		string,
	houseAdLevel2		string,
	houseAdLevel3		string,
	houseStreetName		string,
	houseStreetNum		string,
	houseBuildingNum	string,
	houseRoomNum		string,
	houseLayout			string,
	houseArea			float32,
	price 				uint,
	status				string,
	placerId			string,
	accepterId			string,
	meta 				string){
	
		fmt.Println("pay status ", status);
	ret, err := db.Exec(update_order_by_id, 
		orderType, 
		content,
		houseId,
		houseNation,
		houseAdLevel1,
		houseAdLevel2,
		houseAdLevel3,
		houseStreetName,
		houseStreetNum,
		houseBuildingNum,
		houseRoomNum,
		houseLayout,
		houseArea,
		price, 
		status, 
		placerId, 
		accepterId,
		meta,
		uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update order success !")
	fmt.Println(aff_nums)
}

func UpdateOrderWxPrepayIdByUid(
	wxPrepayId 			string,
	uid 				string){
	fmt.Println("wxPrepayId= ", wxPrepayId)
	fmt.Println("orderId= ", uid)
	sql := fmt.Sprintf(update_order_prepay_id_by_id, wxPrepayId, uid)
	fmt.Println("update sql = ", sql)
	ret, err := db.Exec(sql)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update prepay id for order success !")
	fmt.Println(aff_nums)
}

func UpdateOrderStatusByUid(
	status 			string,
	uid 				string){
		fmt.Println("status= ", status)
		fmt.Println("orderId= ", uid)
	ret, err := db.Exec(update_order_status_by_id, 
		status, 
		uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update status for order success status = ", status)
	fmt.Println(aff_nums)
}

func QueryOrdersByStatusPlacerId(count uint, offset uint, status string, placerId string) []DbOrder{
	
	result := make([]DbOrder, 0)
	rows, err := db.Query(query_orders_by_status_placer, status, placerId, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		result = append(result, *scanOrderItemFromRows(rows))
	}
	return result;
}

func DeleteOrderByUid(uid string){
	//删除数据
	ret, err := db.Exec(update_order_deleted_by_id ,uid)
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		countrySubSql = "ANY(SELECT house_nation FROM order_t)"
	}

	provinceSubSql := "'" + province + "'"
	if(0 == len(province)) {
		provinceSubSql = "ANY(SELECT house_ad_level_1 FROM order_t)"
	}

	citySubSql := "'" + city + "'"
	if(0 == len(city)) {
		citySubSql = "ANY(SELECT house_ad_level_2 FROM order_t)"
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
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
		result = append(result, *scanOrderItemFromRows(rows))
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountByOrderTypeGroup(OrderTypes []string,) uint{

	var temp = "'" + strings.Join(OrderTypes, "','") + "'"
	querySql := fmt.Sprintf(query_order_count_by_order_type_group, temp)
	fmt.Println(querySql)

	var count uint = 0
	err := db.QueryRow(querySql).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderByOrderTypeGroup(OrderTypes []string, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountByOrderTypeGroup(OrderTypes);
	fmt.Println("total = " , total);

	var temp = "'" + strings.Join(OrderTypes, "','") + "'"
	querySql := fmt.Sprintf(query_order_by_order_type_group, temp)
	fmt.Println(querySql)

	rows, err := db.Query(querySql, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		result = append(result, *scanOrderItemFromRows(rows))
	}
	// fmt.Println(result)
	return total, result;
	
}

func QueryOrderTotalCountAll() uint{
	var count uint = 0
	err := db.QueryRow(query_order_count_all).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrders(offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountAll()
	fmt.Println("total = " , total);

	rows, err := db.Query(query_orders, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		result = append(result, *scanOrderItemFromRows(rows))
	}
	// fmt.Println(result)
	return total, result;
	
}

//类型和状态
func QueryOrderTotalCountByOrderTypeGroupStatus(orderTypes []string, status string) uint{

	var temp = "'" + strings.Join(orderTypes, "','") + "'"
	querySql := fmt.Sprintf(query_order_count_by_order_type_group_order_status, temp)
	fmt.Println(querySql)

	var count uint = 0
	err := db.QueryRow(querySql, status).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count;
}

func QueryOrderByOrderTypeGroupStatus(orderTypes []string, status string, offset uint, length uint) (uint, []DbOrder){

	total := QueryOrderTotalCountByOrderTypeGroupStatus(orderTypes, status);
	fmt.Println("total = " , total);

	var temp = "'" + strings.Join(orderTypes, "','") + "'"
	querySql := fmt.Sprintf(query_order_by_order_type_group_order_status, temp)
	fmt.Println(querySql)

	rows, err := db.Query(querySql, status, length, offset)
	defer rows.Close()
	Error.CheckErr(err)
	
	result := make([]DbOrder, 0)
	for rows.Next() {
		result = append(result, *scanOrderItemFromRows(rows))
	}
	// fmt.Println(result)
	return total, result;
	
}


func scanOrderItemFromRows(rows *sql.Rows) *DbOrder{
	var ret = &DbOrder{}
	err := rows.Scan(
		&ret.Uid,
		&ret.OrderType,
		&ret.Content,
		&ret.HouseId,
		&ret.HouseNation, 
		&ret.HouseAdLevel1, 
		&ret.HouseAdLevel2,
		&ret.HouseAdLevel3,
		&ret.HouseStreetName,
		&ret.HouseStreetNum,
		&ret.HouseBuildingNum,
		&ret.HouseRoomNum,  
		&ret.HouseLayout,
		&ret.HouseArea,
		&ret.Price,
		&ret.Status,
		&ret.PlacerId,
		&ret.AccepterId,
		&ret.WxPrepayId,
		&ret.Meta,
		&ret.Deleted,
		&ret.CreateTime)
	Error.CheckErr(err)
	return ret;
}