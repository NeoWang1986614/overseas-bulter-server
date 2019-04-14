package entity

type Error struct {
	Code 			uint	`json:"code"` 
	Message 		string  `json:"message"` 
}

func GetErrForSuccess() *Error{
	return &Error{
		Code: 0,
		Message: "success"}
}