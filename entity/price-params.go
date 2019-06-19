package entity

import (
	storage "overseas-bulter-server/storage"
)

type PriceParams struct {
	Uid 			string 	`json:"uid"`
	ServiceId 		string	`json:"service_id"`
	LayoutId		string	`json:"layout_id"`
	AlgorithmType	string 	`json:"algorithm_type"`
	Params			string	`json:"params"`
	Meta	 		string	`json:"meta"`
}

type PriceParamsRangeQuery struct {
	Offset 			uint	`json:"offset"` 
	Count 			uint  	`json:"count"` 
}

type PriceParamsServiceIdLayoutIdQuery struct {
	ServiceId 		string	`json:"service_id"` 
	LayoutId 		string  `json:"layout_id"` 
}

func ConvertToPriceParamsStorage(enti *PriceParams) *storage.DbPriceParams{
	return &storage.DbPriceParams{
		Uid: enti.Uid,
		ServiceId: enti.ServiceId,
		LayoutId: enti.LayoutId,
		AlgorithmType: enti.AlgorithmType,
		Params: enti.Params,
		Meta: enti.Meta}
}

func ConvertToPriceParamsEntity(obj *storage.DbPriceParams) *PriceParams{
	return &PriceParams{
		Uid: obj.Uid,
		ServiceId: obj.ServiceId,
		LayoutId: obj.LayoutId,
		AlgorithmType: obj.AlgorithmType,
		Params: obj.Params,
		Meta: obj.Meta}
}