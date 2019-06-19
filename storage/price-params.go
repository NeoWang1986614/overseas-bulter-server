package storage

import(
	"fmt"
	"errors"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbPriceParams struct{
	Uid 			string
	ServiceId 		string
	LayoutId		string
	AlgorithmType	string
	Params			string
	Meta	 		string
	Deleted			uint
	CreateTime		string
}

const(
	create_price_params_table_sql = `CREATE TABLE IF NOT EXISTS price_params_t(
		uid VARCHAR(64) NOT NULL unique,
		service_id VARCHAR(64) NULL DEFAULT NULL,
		layout_id VARCHAR(64) NULL DEFAULT NULL,
		algorithm_type VARCHAR(64) NULL DEFAULT NULL,
		params VARCHAR(2048) NULL DEFAULT NULL,
		meta VARCHAR(1024) NULL DEFAULT NULL,
		deleted INT(64) NULL DEFAULT 0,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_price_params = `INSERT INTO price_params_t (
		uid,
		service_id,
		layout_id,
		algorithm_type,
		params,
		meta,
		deleted) VALUE (?,?,?,?,?,?,?)`
	query_price_params = `SELECT * FROM price_params_t WHERE deleted=0 AND uid=?`
	query_price_params_by_service_id_layout_id = `SELECT * FROM price_params_t WHERE deleted=0 AND service_id=? AND layout_id=?`
	query_price_params_by_range = `SELECT * FROM price_params_t WHERE deleted=0 LIMIT ? OFFSET ?`
	update_price_params_by_uid = `UPDATE price_params_t SET 
		service_id=?,
		layout_id=?,
		algorithm_type=?,
		params=?,
		meta=? WHERE uid=?`
	delete_price_params_by_uid = `DELETE FROM price_params_t WHERE uid=?`
)

func CreatePriceParamsTable() {
	sql := create_price_params_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create price params table!");
}

func AddPriceParams(
	serviceId,
	layoutId,
	algorithmType,
	params,
	meta string) error{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_price_params ,uuid, serviceId, layoutId, algorithmType, params, meta, 0);
	// Error.CheckErr(err)
	if(nil != err){
		return errors.New("Error: Add Price Params Error!")
	}
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert price success !")
	fmt.Println(aff_nums);
	return nil;
}

func QueryPriceParams(uid string) *DbPriceParams{
	
	result := &DbPriceParams{}
	rows, err := db.Query(query_price_params, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.ServiceId, 
			&result.LayoutId,
			&result.AlgorithmType,
			&result.Params,  
			&result.Meta,
			&result.Deleted,  
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr price :")
	fmt.Println(result)
	return result
}

func QueryPriceParamsByRange(count, offset uint) []DbPriceParams{

	result := make([]DbPriceParams, 0)
	rows, err := db.Query(query_price_params_by_range, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbPriceParams{}
		err = rows.Scan(
			&item.Uid,
			&item.ServiceId,
			&item.LayoutId,
			&item.AlgorithmType,
			&item.Params,
			&item.Meta,
			&item.Deleted,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func UpdatePriceParams(
	uid				string,
	serviceId 		string,
	layoutId		string,
	algorithmType 	string,
	params			string,
	meta	 		string) {
	//更新数据
	ret, err := db.Exec(update_price_params_by_uid ,serviceId, layoutId, algorithmType, params, meta ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update price params success !")
	fmt.Println(aff_nums);
}

func DeletePriceParams(uid string) {
	//更新数据
	ret, err := db.Exec(delete_price_params_by_uid ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete price params success !")
	fmt.Println(aff_nums);
}

func QueryPriceParamsByServiceIdLayoutId(serviceId, layoutId string) []DbPriceParams{

	result := make([]DbPriceParams, 0)
	rows, err := db.Query(query_price_params_by_service_id_layout_id, serviceId, layoutId)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbPriceParams{}
		err = rows.Scan(
			&item.Uid,
			&item.ServiceId,
			&item.LayoutId,
			&item.AlgorithmType,
			&item.Params,
			&item.Meta,
			&item.Deleted,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}