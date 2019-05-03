package entity

import (
	storage "overseas-bulter-server/storage"
)

type RepairRecord struct {
	Uid 			string	`json:"uid"` 
	OrderId 		string  `json:"order_id"` 
	ReportTime		string	`json:"report_time"`
	RepairTime		string 	`json:"repair_time"`
	CompleteTime	string	`json:"complete_time"`
	Comment			string	`json:"comment"`
	Cost 			float32	`json:"cost"`
	Status	 		string	`json:"status"`
	RelatedImage	string	`json:"related_image"`
	UpdateTime	 	string	`json:"update_time"`
	CreateTIme	 	string	`json:"create_time"` 
}

type RepairRecordQueryByOrderId struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"`
	OrderId			string  `json:"order_id"`
}

type RepairRecordQueryResult struct {
	Total 			uint	`json:"total"` 
	Entities 		[]RepairRecord  	`json:"entities"`
}

func ConvertToRepairRecordStorage(enti *RepairRecord) *storage.DbRepairRecord{
	return &storage.DbRepairRecord{
		Uid: enti.Uid,
		OrderId: enti.OrderId,
		ReportTime: enti.ReportTime,
		RepairTime: enti.RepairTime,
		CompleteTime: enti.CompleteTime,
		Comment: enti.Comment,
		Cost: enti.Cost,
		Status: enti.Status,
		RelatedImage: enti.RelatedImage,
		UpdateTime: enti.UpdateTime,
		CreateTIme: enti.CreateTIme}
}

func ConvertToRepairRecordEntity(obj *storage.DbRepairRecord) *RepairRecord{
	return &RepairRecord{
		Uid: obj.Uid,
		OrderId: obj.OrderId,
		ReportTime: obj.ReportTime,
		RepairTime: obj.RepairTime,
		CompleteTime: obj.CompleteTime,
		Comment: obj.Comment,
		Cost: obj.Cost,
		Status: obj.Status,
		RelatedImage: obj.RelatedImage,
		UpdateTime: obj.UpdateTime,
		CreateTIme: obj.CreateTIme}
}