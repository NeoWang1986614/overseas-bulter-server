package entity

type Feedback struct {
	Uid 				string	`json:"id"`
	OrderId				string	`json:"order_id"`
	AuthorId			string	`json:"author_id"`
	Content 			string  `json:"content"` 
	IsRead 				uint  	`json:"is_read"` 
	Income				float32 `json:"income"`
	Outgoings			float32 `json:"outgoings"`
	AccountingDate		string 	`json:"accounting_date"`
	UpdateTime			string	`json:"update_time"`
	CreateTime			string	`json:"create_time"`
}

type FeedbackQuery struct {
	Offset 			uint	`json:"offset"` 
	Length 			uint  	`json:"length"` 
	OrderId 		string  `json:"order_id"`
	IsRead 			uint  	`json:"is_read"` 
	IsFromBackEnd 	uint  	`json:"is_from_back_end"` 
}