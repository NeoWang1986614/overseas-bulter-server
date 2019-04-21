package entity

type Employee struct{
	Id 				string	`json:"id"`
	UserName		string	`json:"user_name"`
	Password		string	`json:"password"`
	Nickname 		string  `json:"nickname"` 
	AvatarUrl 		string  `json:"avatar_url"` 
	PhoneNumber 	string  `json:"phone_number"`
}

type QueryEmployeeRange struct{
	Offset			uint	`json:"offset"`
	Length			uint	`json:"length"`
}

type EmployeeCheck struct{
	UserName		string	`json:"user_name"`
	Password		string	`json:"password"`
}

type QueryEmployeeResult struct {
	Total				uint 		`json:"total"`
	Entities			[]Employee	`json:"entities"`
}