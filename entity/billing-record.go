package entity

import (
	storage "overseas-bulter-server/storage"
)

type BillingRecord struct {
	Uid 			string	`json:"uid"` 
	OrderId 		string  `json:"order_id"` 
	Income 			string	`json:"income"`
	Outgoings 		string	`json:"outgoings"`
	Balance	 		float32	`json:"balance"`
	Comment	 		string	`json:"comment"`
	TimeRange	 	string	`json:"time_range"`
	AccountingDate	string	`json:"accounting_date"`
	UpdateTime	 	string	`json:"update_time"`
	CreateTIme	 	string	`json:"create_time"` 
}

type BillingRecordQueryByOrderId struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"`
	OrderId			string  `json:"order_id"`
}

type BillingRecordQueryResult struct {
	Total 			uint	`json:"total"` 
	Entities 		[]BillingRecord  	`json:"entities"`
}

func ConvertToBillingRecordStorage(enti *BillingRecord) *storage.DbBillingRecord{
	return &storage.DbBillingRecord{
		Uid: enti.Uid,
		OrderId: enti.OrderId,
		Income: enti.Income,
		Outgoings: enti.Outgoings,
		Balance: enti.Balance,
		Comment: enti.Comment,
		TimeRange: enti.TimeRange,
		AccountingDate: enti.AccountingDate,
		UpdateTime: enti.UpdateTime,
		CreateTIme: enti.CreateTIme}
}

func ConvertToBillingRecordEntity(obj *storage.DbBillingRecord) *BillingRecord{
	return &BillingRecord{
		Uid: obj.Uid,
		OrderId: obj.OrderId,
		Income: obj.Income,
		Outgoings: obj.Outgoings,
		Balance: obj.Balance,
		Comment: obj.Comment,
		TimeRange: obj.TimeRange,
		AccountingDate: obj.AccountingDate,
		UpdateTime: obj.UpdateTime,
		CreateTIme: obj.CreateTIme}
}