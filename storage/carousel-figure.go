package storage

import(
	"fmt"
	"errors"
	_"github.com/go-sql-driver/mysql"
	Error "overseas-bulter-server/error"
	Uuid "overseas-bulter-server/uuid"
)

type DbCarouselFigure struct{
	Uid 		string
	ImageUrl 	string
	Location	string
	Desc		string
	CreateTime	string
}

const(
	create_carousel_figure_table_sql = `CREATE TABLE IF NOT EXISTS carousel_figure_t(
		uid VARCHAR(64) NOT NULL unique,
		image_url VARCHAR(1024) NULL DEFAULT NULL,
		location VARCHAR(64) NULL DEFAULT NULL,
		description VARCHAR(1024) NULL DEFAULT NULL,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY(uid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insert_carousel_figure = `INSERT INTO carousel_figure_t (
		uid,
		image_url,
		location,
		description) VALUE (?,?,?,?)`
	query_carousel_figure = `SELECT * FROM carousel_figure_t WHERE uid=?`
	query_carousel_figure_by_range = `SELECT * FROM carousel_figure_t ORDER BY location LIMIT ? OFFSET ?`
	update_carousel_figure_by_uid = `UPDATE carousel_figure_t SET 
		image_url=?,
		location=?,
		description=? WHERE uid=?`
	delete_carousel_figure_by_uid = `DELETE FROM carousel_figure_t WHERE uid=?`
)

func CreateCarouselFigureTable() {
	sql := create_carousel_figure_table_sql
	smt, err := db.Prepare(sql)
	defer smt.Close()
	Error.CheckErr(err)
	smt.Exec()
	fmt.Println("create carousel figure table!");
}

func AddCarouselFigure(
	imageUrl,
	location,
	desc string) error{
	uuid := Uuid.GenerateNextUuid()
	fmt.Println(uuid)
	//更新数据
	ret, err := db.Exec(insert_carousel_figure,
		uuid,
		imageUrl,
		location,
		desc);
	Error.CheckErr(err)
	if(nil != err){
		return errors.New("Error: Add Carousel Figure Error!")
	}
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("insert carousel figure success !")
	fmt.Println(aff_nums);
	return nil;
}

func QueryCarouselFigure(uid string) *DbCarouselFigure{
	
	result := &DbCarouselFigure{}
	rows, err := db.Query(query_carousel_figure, uid)
	defer rows.Close()
	Error.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(
			&result.Uid, 
			&result.ImageUrl, 
			&result.Location,
			&result.Desc,
			&result.CreateTime)
		Error.CheckErr(err)
	}
	fmt.Println("queyr service primay :")
	fmt.Println(result)
	return result
}

func QueryCarouselFigureByRange(count, offset uint) []DbCarouselFigure{

	result := make([]DbCarouselFigure, 0)
	rows, err := db.Query(query_carousel_figure_by_range, count, offset)
	defer rows.Close()
	Error.CheckErr(err)

	for rows.Next() {
		var item = &DbCarouselFigure{}
		err = rows.Scan(
			&item.Uid,
			&item.ImageUrl, 
			&item.Location,
			&item.Desc,
			&item.CreateTime)
		Error.CheckErr(err)
		result = append(result, *item)
	}
	return result;
}

func UpdateCarouselFigure(
	uid,
	imageUrl,
	location,
	desc string) {
	//更新数据
	ret, err := db.Exec(update_carousel_figure_by_uid,
		imageUrl,
		location,
		desc,
		uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("update carousel figure success !")
	fmt.Println(aff_nums);
}

func DeleteCarouselFigure(uid string) {
	//更新数据
	ret, err := db.Exec(delete_carousel_figure_by_uid ,uid);
	Error.CheckErr(err)
	aff_nums, _ := ret.RowsAffected();
	fmt.Println("delete carousel figure success !")
	fmt.Println(aff_nums)
}