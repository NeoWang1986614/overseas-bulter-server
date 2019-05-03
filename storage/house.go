package storage

import(
	// "database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbHouse struct{//16
	Uid 			string
	Name 			string
	Lat				string
	Lng				string
	AdLevel1		string
	AdLevel2		string
	AdLevel3		string
	Locality		string
	Nation 			string
	StreetName		string
	StreetNum		string
	BuildingNum 	string
	RoomNum			string
	Layout 			string
	OwnerId			string
	CreateTime		string
}

const(
	create_house_table_sql = `CREATE TABLE IF NOT EXISTS house_t(
		uid VARCHAR(64) NOT NULL unique,
		name VARCHAR(64) NULL DEFAULT NULL,
		lat VARCHAR(64) NULL DEFAULT NULL,
		lng VARCHAR(64) NULL DEFAULT NULL,
		ad_level_1 VARCHAR(64) NULL DEFAULT NULL,
		ad_level_2 VARCHAR(64) NULL DEFAULT NULL,
		ad_level_3 VARCHAR(64) NULL DEFAULT NULL,
		locality VARCHAR(64) NULL DEFAULT NULL,
		nation VARCHAR(64) NULL DEFAULT NULL,
		street_name VARCHAR(64) NULL DEFAULT NULL,
		street_num VARCHAR(64) NULL DEFAULT NULL,
		building_num VARCHAR(64) NULL DEFAULT NULL,
		room_num VARCHAR(64) NULL DEFAULT NULL,
		layout VARCHAR(64) NULL DEFAULT NULL,
		owner_id VARCHAR(64) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_house = `INSERT INTO house_t (
						uid,
						name,
						lat,
						lng,
						ad_level_1,
						ad_level_2,
						ad_level_3,
						locality,
						nation,
						street_name,
						street_num,
						building_num,
						room_num,
						layout,
						owner_id) VALUE (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	query_house_by_range = `SELECT * FROM house_t WHERE owner_id=? LIMIT ? OFFSET ?`
	query_house = `SELECT * FROM house_t WHERE uid=?`
	update_house_by_uid = `UPDATE house_t SET 
	name=?,
	lat=?, 
	lng=?,
	ad_level_1=?,
	ad_level_2=?,
	ad_level_3=?,
	locality=?,
	nation=?,
	street_name=?,
	street_num=?,
	building_num=?,
	room_num=?,
	layout=?,
	owner_id=? WHERE uid=?`
	delete_house_by_uid = `DELETE FROM house_t WHERE uid=?`
)

func CreateHouseTable() {
	sql := create_house_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create house table!");
	// arr := QueryHousesByRange(5, 0)
}

func QueryHouse(id string) *DbHouse{
	
	result := &DbHouse{}
	rows, err := db.Query(query_house, id)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.Name, 
			&result.Lat,  
			&result.Lng, 
			&result.AdLevel1, 
			&result.AdLevel2, 
			&result.AdLevel3, 
			&result.Locality, 
			&result.Nation, 
			&result.StreetName,
			&result.StreetNum,
			&result.BuildingNum,
			&result.RoomNum, 
			&result.Layout,
			&result.OwnerId,
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr house :")
	fmt.Println(result)
	return result
}

func AddHouse(
	name string,
	lat string,
	lng string,
	adLevel1 string,
	adLevel2 string,
	adLevel3 string,
	locality string,
	nation string,
	streetName string,
	streetNum string,
	buildingNum string,
	roomNum string,
	layout string,
	ownerId string) string{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_house,
		uuid, 
		name, 
		lat, 
		lng,
		adLevel1,
		adLevel2,
		adLevel3,
		locality,
		nation,
		streetName,
		streetNum,
		buildingNum,
		roomNum,
		layout,
		ownerId);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert house success !")
	fmt.Println(aff_nums);
	return uuid;
}

func UpdateHouse(
	uid string,
	name string,
	lat string,
	lng string,
	adLevel1 string,
	adLevel2 string,
	adLevel3 string,
	locality string,
	nation string,
	streetName string,
	streetNum string,
	buildingNum string,
	roomNum string,
	layout string,
	ownerId string){
	
	ret, err := db.Exec(update_house_by_uid, 
		name,
		lat, 
		lng,
		adLevel1,
		adLevel2,
		adLevel3,
		locality,
		nation,
		streetName,
		streetNum,
		buildingNum,
		roomNum, 
		layout, 
		ownerId, 
		uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update house success !")
	fmt.Println(aff_nums)
}

func QueryHouses(ownerId string, count uint, offset uint) []DbHouse{
	
	result := make([]DbHouse, 0)
	rows, err := db.Query(query_house_by_range, ownerId, count, offset)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		var item = &DbHouse{}
		err = rows.Scan(
			&item.Uid,
			&item.Name,
			&item.Lat,
			&item.Lng,
			&item.AdLevel1,
			&item.AdLevel2,
			&item.AdLevel3,
			&item.Locality,
			&item.Nation,
			&item.StreetName,
			&item.StreetNum,
			&item.BuildingNum,
			&item.RoomNum,
			&item.Layout,
			&item.OwnerId,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func DeleteHouseByUid(uid string){
	//删除数据
	ret, err := db.Exec(delete_house_by_uid ,uid)
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete house success !")
	fmt.Println(aff_nums);
}