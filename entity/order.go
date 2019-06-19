package entity

import (
	storage "overseas-bulter-server/storage"
)

type Order struct {
	Uid 				string	`json:"id"`
	Type				string	`json:"type"`
	Content 			string  `json:"content"`
	HouseId 			string  `json:"house_id"` 
	HouseNation 		string	`json:"house_nation"`
	HouseAdLevel1 		string	`json:"house_ad_level_1"`
	HouseAdLevel2 		string	`json:"house_ad_level_2"`
	HouseAdLevel3 		string	`json:"house_ad_level_3"`
	HouseStreetName 	string	`json:"house_street_name"`
	HouseStreetNum	 	string	`json:"house_street_num"`
	HouseBuildingNum 	string	`json:"house_building_num"`
	HouseRoomNum 		string	`json:"house_room_num"`
	HouseLayout 		string	`json:"house_layout"`
	HouseArea 			float32	`json:"house_area"`
	Price				uint	`json:"price"`
	Status				string  `json:"status"`
	PlacerId			string	`json:"placer_id"`
	AccepterId			string	`json:"accepter_id"`
	WxPrepayId			string	`json:"wx_prepay_id"`
	Meta			string	`json:"meta"`
	CreateTime			string	`json:"create_time"`
}

type OrderSearchByStatusPalcerId struct {
	Offset 			uint	`json:"offset"` 
	Length 			uint  	`json:"length"` 
	Status 			string  `json:"status"` 
	PlacerId 		string  `json:"placer_id"` 
}

type AddOrderResult struct {
	Id 			string	`json:"id"` 
}

type OrderQueryByIdCardNumber struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	PlacerIdCardNumber	string 	`json:"id_card_number"` 
}

type OrderQueryByPhoneNumber struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	PlacerPhoneNumber	string 	`json:"phone_number"` 
}

type OrderQueryByRealName struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	RealName			string 	`json:"real_name"` 
}

type OrderQueryBeforeTime struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	Time				string 	`json:"time"` 
}

type OrderQueryAfterTime struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	Time				string 	`json:"time"` 
}

type OrderQueryRangeTime struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	FromTime			string 	`json:"from_time"` 
	ToTime				string 	`json:"to_time"` 
}

type OrderQueryByStatusGroup struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	StatusGroup			[]string `json:"status_group"` 
}

type OrderQueryByAddress struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	Country				string 	`json:"country"` 
	Province			string 	`json:"province"` 
	City				string 	`json:"city"` 
}

type OrderQueryByLayoutGroup struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	LayoutGroup			[]string `json:"layout_group"` 
}

type OrderQueryBelowPrice struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	Price				uint 	`json:"price"` 
}

type OrderQueryAbovePrice struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	Price				uint 	`json:"price"` 
}

type OrderQueryRangePrice struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	FromPrice			uint 	`json:"from_price"` 
	ToPrice				uint 	`json:"to_price"` 
}

type OrderQueryByOrderTypeGroup struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
	TypeGroup			[]string `json:"type_group"` 
}

type OrderQueryAll struct {
	Offset 				uint	`json:"offset"` 
	Length 				uint  	`json:"length"`
}

type OrderQueryResult struct {
	Total				uint 	`json:"total"`
	Entities			[]Order	`json:"entities"`
}

func ConvertToOrderStorage(enti *Order) *storage.DbOrder{
	return &storage.DbOrder{
		Uid: enti.Uid,
		OrderType: enti.Type,
		Content: enti.Content,
		HouseId: enti.HouseId,
		HouseNation: enti.HouseNation,
		HouseAdLevel1: enti.HouseAdLevel1,
		HouseAdLevel2: enti.HouseAdLevel2,
		HouseAdLevel3: enti.HouseAdLevel3,
		HouseStreetName: enti.HouseStreetName,
		HouseStreetNum: enti.HouseStreetNum,
		HouseBuildingNum: enti.HouseBuildingNum,
		HouseRoomNum: enti.HouseRoomNum,
		HouseLayout: enti.HouseLayout,
		HouseArea: enti.HouseArea,
		Price: enti.Price,
		Status: enti.Status,
		PlacerId: enti.PlacerId,
		AccepterId: enti.AccepterId,
		WxPrepayId: enti.WxPrepayId,
		Meta: enti.Meta,
		CreateTime: enti.CreateTime}
}

func ConvertToOrderEntity(obj *storage.DbOrder) *Order{
	return &Order{
		Uid: obj.Uid,
		Type: obj.OrderType,
		Content: obj.Content,
		HouseId: obj.HouseId,
		HouseNation: obj.HouseNation,
		HouseAdLevel1: obj.HouseAdLevel1,
		HouseAdLevel2: obj.HouseAdLevel2,
		HouseAdLevel3: obj.HouseAdLevel3,
		HouseStreetName: obj.HouseStreetName,
		HouseStreetNum: obj.HouseStreetNum,
		HouseBuildingNum: obj.HouseBuildingNum,
		HouseRoomNum: obj.HouseRoomNum,
		HouseLayout: obj.HouseLayout,
		HouseArea: obj.HouseArea,
		Price: obj.Price,
		Status: obj.Status,
		PlacerId: obj.PlacerId,
		AccepterId: obj.AccepterId,
		WxPrepayId: obj.WxPrepayId,
		Meta: obj.Meta,
		CreateTime: obj.CreateTime}
}