package entity

type LoginRequest struct {
	Code 			string	`json:"code"` 
	AppId 			string  `json:"app_id"`
	AppSecret 		string  `json:"app_secret"`
}