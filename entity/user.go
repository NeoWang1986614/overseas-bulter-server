package entity

type User struct {
	Uid          	string  `json:"uid"`
	Name          	string  `json:"name"`
	PhoneNumber		string 	`json:"phone_number"`
	IdCardNumber    string  `json:"id_card_number"`
}

type QueryUser struct {
	Uid          	string  `json:"uid"`
}