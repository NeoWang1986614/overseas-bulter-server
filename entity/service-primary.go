package entity

import (
	storage "overseas-bulter-server/storage"
)

type ServicePrimary struct {
	Uid 				string	`json:"uid"`
	Value				string	`json:"value"`
	Title				string	`json:"title"`
	IconUrl				string	`json:"icon_url"`
	Location			string	`json:"location"`
	Content 			string  `json:"content"`
	BasePrice 			string  `json:"base_price"`
	Meta				string	`json:"meta"`
}

type ServicePrimaryRangeQuery struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"` 
}

func ConvertToServicePrimaryStorage(enti *ServicePrimary) *storage.DbServicePrimary{
	return &storage.DbServicePrimary{
		Uid: enti.Uid,
		Value: enti.Value,
		Title: enti.Title,
		IconUrl: enti.IconUrl,
		Location: enti.Location,
		Content: enti.Content,
		BasePrice: enti.BasePrice,
		Meta: enti.Meta}
}

func ConvertToServicePrimaryEntity(obj *storage.DbServicePrimary) *ServicePrimary{
	return &ServicePrimary{
		Uid: obj.Uid,
		Value: obj.Value,
		Title: obj.Title,
		IconUrl: obj.IconUrl,
		Location: obj.Location,
		Content: obj.Content,
		BasePrice: obj.BasePrice,
		Meta: obj.Meta}
}