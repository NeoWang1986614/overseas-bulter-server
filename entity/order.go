package entity

type Order struct {
	Uid 				string	`json:"id"`
	Type				string	`json:"type"`
	Content 			string  `json:"content"` 
	HouseCountry 		string	`json:"house_country"`
	HouseProvince 		string	`json:"house_province"`
	HouseCity 			string	`json:"house_city"`
	HouseAddress 		string	`json:"house_address"`
	HouseLayout 		string	`json:"house_layout"`
	Price				uint	`json:"price"`
	Status				string  `json:"status"`
	PlacerId			string	`json:"placer_id"`
	AccepterId			string	`json:"accepter_id"`
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