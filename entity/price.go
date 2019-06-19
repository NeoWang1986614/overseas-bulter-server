package entity

type ComputePriceParam struct {
	ServiceId 		string	`json:"service_id"`
	LayoutId		string	`json:"layout_id"`
	Param			string 	`json:"param"`
}

type AlgorithmType0Param struct {
	Area 			float32	`json:"area"`
}

type AlgorithmType0Config struct {
	BaseArea 		float32	`json:"baseArea"`
	BaseFee 		float32	`json:"baseFee"`
	StepArea 		float32	`json:"stepArea"`
	StepFee 		float32	`json:"stepFee"`
	Discount 		float32	`json:"discount"`
}

type AlgorithmType1Param struct {
	Year 			uint	`json:"year"`
}

type AlgorithmType1Config struct {
	FeePerYear 		float32		`json:"feePerYear"`
	TimeNodes 		[]uint		`json:"timeNodes"`
	Discounts 		[]float32	`json:"discounts"`
}

type ComputePriceResult struct {
	OldPrice 		float32	`json:"old_price"`
	NewPrice		float32	`json:"new_price"`
	Discount		float32 `json:"discount"`
}