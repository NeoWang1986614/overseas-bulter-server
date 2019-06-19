package entity

import (
	storage "overseas-bulter-server/storage"
)

type CarouselFigure struct {
	Uid 				string	`json:"uid"`
	ImageUrl			string	`json:"image_url"`
	Location			string	`json:"location"`
	Desc				string	`json:"desc"`
}

type CarouselFigureRangeQuery struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"` 
}

func ConvertToCarouselFigureStorage(enti *CarouselFigure) *storage.DbCarouselFigure{
	return &storage.DbCarouselFigure{
		Uid: enti.Uid,
		ImageUrl: enti.ImageUrl,
		Location: enti.Location,
		Desc: enti.Desc}
}

func ConvertToCarouselFigureEntity(obj *storage.DbCarouselFigure) *CarouselFigure{
	return &CarouselFigure{
		Uid: obj.Uid,
		ImageUrl: obj.ImageUrl,
		Location: obj.Location,
		Desc: obj.Desc}
}