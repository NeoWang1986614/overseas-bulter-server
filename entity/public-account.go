package entity

type PublicAccountMaterialQuery struct {
	Type 			string	`json:"type"` 
	Offset 			uint  	`json:"offset"`
	Count			uint  	`json:"count"`
}