package entity

import (
	storage "overseas-bulter-server/storage"
)

type HouseDeal struct {
	Uid 			string	`json:"uid"` 
	DealType 		string  `json:"deal_type"`
	Source 			string  `json:"source"`
	HouseId 		string  `json:"house_id"`
	Decoration	 	string	`json:"decoration"`
	Cost	 		string	`json:"cost"`
	Linkman	 		string	`json:"linkman"`
	ContactNum	 	string	`json:"contact_num"`
	Mail	 		string	`json:"mail"`
	Weixin	 		string	`json:"weixin"`
	Image	 		string	`json:"image"`
	Note	 		string	`json:"note"`
	Creator	 		string	`json:"creator"`
	Meta	 		string	`json:"meta"`
	CreateTime 		string  `json:"create_time"`
}

type HouseDealRangeQuery struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"`
}

type HouseDealQueryByDealType struct {
	DealType 		string	`json:"deal_type"`
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"`
}

type HouseDealQueryResult struct {
	Total 			uint			`json:"total"` 
	Entities 		[]HouseDeal  	`json:"entities"`
}


func ConvertToHouseDealStorage(enti *HouseDeal) *storage.DbHouseDeal{
	return &storage.DbHouseDeal{
		Uid: enti.Uid,
		DealType: enti.DealType,
		Source: enti.Source,
		HouseId: enti.HouseId,
		Decoration: enti.Decoration,
		Cost: enti.Cost,
		Linkman: enti.Linkman,
		ContactNum: enti.ContactNum,
		Mail: enti.Mail,
		Weixin: enti.Weixin,
		Image: enti.Image,
		Note: enti.Note,
		Creator: enti.Creator,
		Meta: enti.Meta,
		CreateTime: enti.CreateTime}
}

func ConvertToHouseDealEntity(obj *storage.DbHouseDeal) *HouseDeal{
	return &HouseDeal{
		Uid: obj.Uid,
		DealType: obj.DealType,
		Source: obj.Source,
		HouseId: obj.HouseId,
		Decoration: obj.Decoration,
		Cost: obj.Cost,
		Linkman: obj.Linkman,
		ContactNum: obj.ContactNum,
		Mail: obj.Mail,
		Weixin: obj.Weixin,
		Image: obj.Image,
		Note: obj.Note,
		Creator: obj.Creator,
		Meta: obj.Meta,
		CreateTime: obj.CreateTime}
}