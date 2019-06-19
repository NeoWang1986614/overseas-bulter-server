package storage

import(
	"fmt"
	"errors"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbHouseDeal struct{
	Uid 			string
	DealType		string
	Source			string
	HouseId			string
	// Property 		string
	// Area	 		string
	// Address			string
	Decoration		string
	Cost			string
	Linkman			string
	ContactNum		string
	Mail			string
	Weixin	 		string
	Image			string
	Note			string
	Creator			string
	Meta 			string
	Deleted			uint
	CreateTime		string
}

const(
	create_house_deal_table_sql = `CREATE TABLE IF NOT EXISTS house_deal_t(
		uid VARCHAR(64) NOT NULL unique,
		deal_type VARCHAR(64) NULL DEFAULT NULL,
		source VARCHAR(64) NULL DEFAULT NULL,
		house_id VARCHAR(64) NULL DEFAULT NULL,
		decoration VARCHAR(64) NULL DEFAULT NULL,
		cost VARCHAR(64) NULL DEFAULT NULL,
		linkman VARCHAR(64) NULL DEFAULT NULL,
		contact_num VARCHAR(64) NULL DEFAULT NULL,
		mail VARCHAR(64) NULL DEFAULT NULL,
		weixin VARCHAR(64) NULL DEFAULT NULL,
		image VARCHAR(2048) NULL DEFAULT NULL,
		note VARCHAR(2048) NULL DEFAULT NULL,
		creator VARCHAR(64) NULL DEFAULT NULL,
		meta VARCHAR(1024) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT 0,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_house_deal = `INSERT INTO house_deal_t (
		uid,
		deal_type,
		source,
		house_id,
		decoration,
		cost,
		linkman,
		contact_num,
		mail,
		weixin,
		image,
		note,
		creator,
		meta,
		deleted) VALUE (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	query_house_deal = `SELECT * FROM house_deal_t WHERE deleted=0 AND uid=?`
	
	query_house_deal_by_range = `SELECT * FROM house_deal_t WHERE deleted=0 ORDER BY create_time LIMIT ? OFFSET ?`
	query_house_deal_by_deal_type = `SELECT * FROM house_deal_t WHERE deal_type=? AND deleted=0 ORDER BY source LIMIT ? OFFSET ?`
	query_house_deal_count_by_deal_type = `SELECT COUNT(*) FROM house_deal_t WHERE deal_type=? AND deleted=0`

	update_house_deal_by_uid = `UPDATE house_deal_t SET 
		deal_type=?,
		source=?,
		house_id=?,
		decoration=?,
		cost=? 
		linkman=?,
		contact_num=?,
		mail=?,
		weixin=?,
		image=?,
		note=?,
		creator=?,
		meta=? WHERE uid=?`
	update_house_deal_deleted_by_uid = `UPDATE house_deal_t SET deleted=1 WHERE uid=?`
	delete_house_deal_deleted_by_uid = `DELETE FROM house_deal_t WHERE uid=?`
)

func CreateHouseDealTable() {
	sql := create_house_deal_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create house deal table!");
}

func AddHouseDeal(
	dealType,
	source,
	house_id,
	decoration,
	cost,
	linkman,
	contactNum,
	mail,
	weixin,
	image,
	note,
	creator,
	meta string) error{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_house_deal,
		uuid,
		dealType,
		source,
		house_id,
		decoration,
		cost,
		linkman,
		contactNum,
		mail,
		weixin,
		image,
		note,
		creator,
		meta,
		0);
	Error.CheckErr(err)
	if(nil != err){
		return errors.New("Error: Add House Deal Error!")
	}
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert house deal success !")
	fmt.Println(aff_nums);
	return nil;
}

func QueryHouseDeal(uid string) *DbHouseDeal{
	
	result := &DbHouseDeal{}
	rows, err := db.Query(query_house_deal, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.DealType,
			&result.Source,
			&result.HouseId,
			&result.Decoration,  
			&result.Cost,
			&result.Linkman,
			&result.ContactNum,
			&result.Mail,
			&result.Weixin,
			&result.Image,
			&result.Note,
			&result.Creator,
			&result.Meta,
			&result.Deleted,  
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr house deal :")
	fmt.Println(result)
	return result
}

func QueryHouseDealByRange(count, offset uint) []DbHouseDeal{

	result := make([]DbHouseDeal, 0)
	rows, err := db.Query(query_house_deal_by_range, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbHouseDeal{}
		err = rows.Scan(
			&item.Uid,
			&item.DealType,
			&item.Source,
			&item.HouseId,
			&item.Decoration,  
			&item.Cost,
			&item.Linkman,
			&item.ContactNum,
			&item.Mail,
			&item.Weixin,
			&item.Image,
			&item.Note,
			&item.Creator,
			&item.Meta,
			&item.Deleted,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func QueryHouseDealCountByDealType(dealType string) uint{
	var count uint = 0
	err := db.QueryRow(query_house_deal_count_by_deal_type, dealType).Scan(&count)
	Error.CheckErr(err)
	fmt.Println(count)
	return count
}

func QueryHouseDealByDealType(dealType string, count, offset uint) (uint,[]DbHouseDeal){

	total := QueryHouseDealCountByDealType(dealType)
	fmt.Println("total = " , total);

	result := make([]DbHouseDeal, 0)
	rows, err := db.Query(query_house_deal_by_deal_type, dealType, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbHouseDeal{}
		err = rows.Scan(
			&item.Uid,
			&item.DealType,
			&item.Source,
			&item.HouseId,
			&item.Decoration,  
			&item.Cost,
			&item.Linkman,
			&item.ContactNum,
			&item.Mail,
			&item.Weixin,
			&item.Image,
			&item.Note,
			&item.Creator,
			&item.Meta,
			&item.Deleted,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return total, result;
}

func UpdateHouseDeal(
	uid,
	dealType,
	source,
	houseId,
	decoration,
	cost,
	linkman,
	contactNum,
	mail,
	weixin,
	image,
	note,
	creator,
	meta string) {
	//更新数据
	ret, err := db.Exec(update_house_deal_by_uid,
		dealType,
		source,
		houseId,
		decoration,
		cost,
		linkman,
		contactNum,
		mail,
		weixin,
		image,
		note,
		creator,
		meta,
		uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update house deal success !")
	fmt.Println(aff_nums);
}

func DeleteHouseDeal(uid string) {
	//更新数据
	ret, err := db.Exec(update_house_deal_deleted_by_uid ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete house deal success !")
	fmt.Println(aff_nums);
}