package entity

type AddShareRequest struct {
	FromUid			string	`json:"uid"` 
	EncryptedData 	string  `json:"encrypted_data"`
	Iv 				string  `json:"iv"`
}

type WaterMarkType struct {
	Timestamp 	uint `json:"timestamp"`
	AppId 		string `json:"appId"`
}

type ShareInfoOpenData struct {
	OpenGId 	string `json:"openGId"`
	WaterMark  	WaterMarkType
}