package entity

import (
	storage "overseas-bulter-server/storage"
)

type InspectRecord struct {
	Uid 			string	`json:"uid"` 
	OrderId 		string  `json:"order_id"` 
	InspectDate		string	`json:"inspect_date"`
	Inspector		string 	`json:"inspector"`
	Comment	 		string	`json:"comment"`
	Config 			string	`json:"config"`
	Area	 		string	`json:"area"`
	UpdateTime	 	string	`json:"update_time"`
	CreateTIme	 	string	`json:"create_time"` 
}

type InspectRecordQueryByOrderId struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"`
	OrderId			string  `json:"order_id"`
}

type InspectRecordQueryResult struct {
	Total 			uint	`json:"total"` 
	Entities 		[]InspectRecord  	`json:"entities"`
}

func ConvertToInspectRecordStorage(enti *InspectRecord) *storage.DbInspectRecord{
	return &storage.DbInspectRecord{
		Uid: enti.Uid,
		OrderId: enti.OrderId,
		InspectDate: enti.InspectDate,
		Inspector: enti.Inspector,
		Comment: enti.Comment,
		Config: enti.Config,
		Area: enti.Area,
		UpdateTime: enti.UpdateTime,
		CreateTIme: enti.CreateTIme}
}

func ConvertToInspectRecordEntity(obj *storage.DbInspectRecord) *InspectRecord{
	return &InspectRecord{
		Uid: obj.Uid,
		OrderId: obj.OrderId,
		InspectDate: obj.InspectDate,
		Inspector: obj.Inspector,
		Comment: obj.Comment,
		Config: obj.Config,
		Area: obj.Area,
		UpdateTime: obj.UpdateTime,
		CreateTIme: obj.CreateTIme}
}