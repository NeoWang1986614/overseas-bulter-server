package entity

import (
	storage "overseas-bulter-server/storage"
)

type Layout struct {
	Uid 				string	`json:"uid"`
	Value				string	`json:"value"`
	Title				string	`json:"title"`
	Location			string  `json:"location"`
	Content 			string  `json:"content"` 
	Meta				string	`json:"meta"`
}

type LayoutRangeQuery struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"` 
}

func ConvertToLayoutStorage(enti *Layout) *storage.DbLayout{
	return &storage.DbLayout{
		Uid: enti.Uid,
		Value: enti.Value,
		Title: enti.Title,
		Location: enti.Location,
		Content: enti.Content,
		Meta: enti.Meta}
}

func ConvertToLayoutEntity(obj *storage.DbLayout) *Layout{
	return &Layout{
		Uid: obj.Uid,
		Value: obj.Value,
		Title: obj.Title,
		Location: obj.Location,
		Content: obj.Content,
		Meta: obj.Meta}
}