package entity

type PublicAccountMaterialQuery struct {
	Type 			string	`json:"type"` 
	Offset 			uint  	`json:"offset"`
	Count			uint  	`json:"count"`
}

// func ConvertToHouseStorage(enti *House) *storage.DbHouse{
// 	return &storage.DbHouse{
// 		Uid: enti.Uid,
// 		Name: enti.Name,
// 		Lat: enti.Lat,
// 		Lng: enti.Lng,
// 		AdLevel1: enti.AdLevel1,
// 		AdLevel2: enti.AdLevel2,
// 		AdLevel3: enti.AdLevel3,
// 		Locality: enti.Locality,
// 		Nation: enti.Nation,
// 		StreetName: enti.StreetName,
// 		StreetNum: enti.StreetNum,
// 		BuildingNum: enti.BuildingNum,
// 		RoomNum: enti.RoomNum,
// 		Layout: enti.Layout,
// 		OwnerId: enti.OwnerId}
// }

// func ConvertToHouseEntity(obj *storage.DbHouse) *House{
// 	return &House{
// 		Uid: obj.Uid,
// 		Name: obj.Name,
// 		Lat: obj.Lat,
// 		Lng: obj.Lng,
// 		AdLevel1: obj.AdLevel1,
// 		AdLevel2: obj.AdLevel2,
// 		AdLevel3: obj.AdLevel3,
// 		Locality: obj.Locality,
// 		Nation: obj.Nation,
// 		StreetName: obj.StreetName,
// 		StreetNum: obj.StreetNum,
// 		BuildingNum: obj.BuildingNum,
// 		RoomNum: obj.RoomNum,
// 		Layout: obj.Layout,
// 		OwnerId: obj.OwnerId}
// }