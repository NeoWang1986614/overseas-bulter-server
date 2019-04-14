package entity

type Service struct {
	Uid 			string	`json:"id"`
	Type			string	`json:"type"`
	Layout			string	`json:"layout"`
	Content			string	`json:"content"`
	Price 			uint  	`json:"price"` 
	CreateTime 		string  `json:"create_time"` 
}

type ServiceQuery struct {
	Type 			string	`json:"type"` 
}