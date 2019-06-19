package entity

type RentHouse struct {
	OrderId 		string 	`json:"order_id"`
	HouseImage 		string	`json:"house_image"`
	HouseLayout		string  `json:"house_layout"`
	HouseAdLevel2	string 	`json:"house_ad_level_2"`
	HouseAdLevel3	string 	`json:"house_ad_level_3"`
	HouseName		string	`json:"house_name"`
	OrderMeta		string	`json:"order_meta"`
}

type QuerySearchRentHouse struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"` 
}

type RentHouseQueryResult struct {
	Total				uint 	`json:"total"`
	Entities			[]RentHouse	`json:"entities"`
}