package entity

type Prepayment struct {
	UserId 		string 	`json:"user_id"`
	OrderId		string		`json:"order_id"`
}

type PrepaymentResult struct {
	NonceStr		string		`json:"nonce_str"`
	PrepayId		string 		`json:"prepay_id"`
	SignType		string 		`json:"sign_type"`
	Sign			string 		`json:"sign"`
	Timestamp		int64		`json:"timestamp"`
	AppId			string		`json:"app_id"`
}
