package entity

import (
	storage "overseas-bulter-server/storage"
)

type House struct {
	Uid 			string	`json:"uid"` 
	Name 			string  `json:"name"`
	Property 		string  `json:"property"` 
	Lat 			string	`json:"lat"`
	Lng 			string	`json:"lng"`
	AdLevel1	 	string	`json:"ad_level_1"`
	AdLevel2	 	string	`json:"ad_level_2"`
	AdLevel3	 	string	`json:"ad_level_3"`
	Locality	 	string	`json:"locality"`
	Nation	 		string	`json:"nation"`
	StreetName	 	string	`json:"street_name"`
	StreetNum	 	string	`json:"street_num"`
	BuildingNum	 	string	`json:"building_num"`
	RoomNum	 		string	`json:"room_num"`
	Layout 			string 	`json:"layout"`
	Area 			float32 `json:"area"`
	OwnerId 		string 	`json:"owner_id"`
	Status 			string 	`json:"status"`
	Meta 			string  `json:"meta"`
}

type HouseQueryByOwnerId struct {
	Offset 			uint	`json:"offset"` 
	Length 			uint  	`json:"length"`
	OwnerId			string  `json:"owner_id"`
}

type HouseQueryByUidGroup struct {
	Uids 			[]string	`json:"uids"` 
}

type AddHouseResult struct {
	Uid          	string  `json:"uid"`
}

type UpdateHouseResult struct {
	Uid          	string  `json:"uid"`
}

func ConvertToHouseStorage(enti *House) *storage.DbHouse{
	return &storage.DbHouse{
		Uid: enti.Uid,
		Name: enti.Name,
		Property: enti.Property,
		Lat: enti.Lat,
		Lng: enti.Lng,
		AdLevel1: enti.AdLevel1,
		AdLevel2: enti.AdLevel2,
		AdLevel3: enti.AdLevel3,
		Locality: enti.Locality,
		Nation: enti.Nation,
		StreetName: enti.StreetName,
		StreetNum: enti.StreetNum,
		BuildingNum: enti.BuildingNum,
		RoomNum: enti.RoomNum,
		Layout: enti.Layout,
		Area: enti.Area,
		OwnerId: enti.OwnerId,
		Status: enti.Status,
		Meta: enti.Meta}
}

func ConvertToHouseEntity(obj *storage.DbHouse) *House{
	return &House{
		Uid: obj.Uid,
		Name: obj.Name,
		Property: obj.Property,
		Lat: obj.Lat,
		Lng: obj.Lng,
		AdLevel1: obj.AdLevel1,
		AdLevel2: obj.AdLevel2,
		AdLevel3: obj.AdLevel3,
		Locality: obj.Locality,
		Nation: obj.Nation,
		StreetName: obj.StreetName,
		StreetNum: obj.StreetNum,
		BuildingNum: obj.BuildingNum,
		RoomNum: obj.RoomNum,
		Layout: obj.Layout,
		Area: obj.Area,
		OwnerId: obj.OwnerId,
		Status: obj.Status,
		Meta: obj.Meta}
}