package entity

type Case struct {
	Uid 			string	`json:"uid"` 
	Title 			string  `json:"title"` 
	ImageUrl 		string	`json:"image_url"`
	Content 		string	`json:"content"`
	Price	 		uint	`json:"price"`
	Level 			uint	`json:"level"`
}

type CaseSearchByLevel struct {
	Offset 			uint	`json:"offset"` 
	Length 			uint  	`json:"length"` 
	Level 			uint  	`json:"level"` 
}

type AddCaseResult struct {
	Uid 			string	`json:"uid"` 
}
