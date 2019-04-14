package entity

type House struct {
	Uid 			string	`json:"uid"` 
	Name 			string  `json:"name"` 
	Country 		string	`json:"country"`
	Province 		string	`json:"province"`
	City	 		string	`json:"city"`
	Address 		string	`json:"address"`
	Layout 			string 	`json:"layout"`  
	OwnerId 		string 	`json:"owner_id"` 
}

type HouseSearch struct {
	Offset 			uint	`json:"offset"` 
	Length 			uint  	`json:"length"` 
}

type AddHouseResult struct {
	Uid          	string  `json:"uid"`
}

type UpdateHouseResult struct {
	Uid          	string  `json:"uid"`
}
