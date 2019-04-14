package entity

type Protocol struct {
	Id 				string	`json:"id"` 
	Title 			string  `json:"title"` 
	AssetID 		string	`json:"asset_id"`
	Price			uint	`json:"price"`
	Content 		string 	`json:"content"`
}
