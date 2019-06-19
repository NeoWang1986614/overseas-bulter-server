package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	"math"
	// // wx "overseas-bulter-server/wx"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
	storage "overseas-bulter-server/storage"
	// entity "overseas-bulter-server/entity"
)

func PriceHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postPriceHandler(w, r)
		break;
	case "PUT":
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func postPriceHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.ComputePriceParam{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	priceParams := storage.QueryPriceParamsByServiceIdLayoutId(requestBody.ServiceId, requestBody.LayoutId)
	ret := &entity.ComputePriceResult{}
	
	if(0 != len(priceParams)){
		curPriceParams := priceParams[0]
		fmt.Print(curPriceParams)
	
		if("0" == curPriceParams.AlgorithmType){
	
			params := &entity.AlgorithmType0Param{}
			err = json.Unmarshal([]byte(requestBody.Param), params)
			Error.CheckErr(err)
			fmt.Print(params)
	
			config := &entity.AlgorithmType0Config{}
			err = json.Unmarshal([]byte(curPriceParams.Params), config)
			Error.CheckErr(err)
			fmt.Print(config)
	
			ret = computeAlgorithmType0(params, config)
	
		}else if("1" == curPriceParams.AlgorithmType){
	
			params := &entity.AlgorithmType1Param{}
			err = json.Unmarshal([]byte(requestBody.Param), params)
			Error.CheckErr(err)
			fmt.Print(params)
	
			config := &entity.AlgorithmType1Config{}
			err = json.Unmarshal([]byte(curPriceParams.Params), config)
			Error.CheckErr(err)
			fmt.Print(config)
	
			ret = computeAlgorithmType1(params, config)
		}
	}
	
	
	rsp, err := json.Marshal(ret)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func computeAlgorithmType0(param *entity.AlgorithmType0Param, config *entity.AlgorithmType0Config) *entity.ComputePriceResult{
	fmt.Println("param.Area=",param.Area)
	fmt.Println("config.BaseArea=",config.BaseArea)
	fmt.Println("config.StepArea=",config.StepArea)
	fmt.Println("config.BaseFee=",config.BaseFee)
	fmt.Println("config.StepFee=",config.StepFee)
	fmt.Println("config.Discount=",config.Discount)

	ret := &entity.ComputePriceResult{}

	if(param.Area < config.BaseArea){
		ret.OldPrice = config.BaseFee
	}else{
		ret.OldPrice = config.BaseFee + config.StepFee * (float32)(math.Ceil( float64((param.Area - config.BaseArea)/config.StepArea)))
	}
	ret.NewPrice = ret.OldPrice * config.Discount
	ret.Discount = config.Discount
	return ret
}

func computeAlgorithmType1(param *entity.AlgorithmType1Param, config *entity.AlgorithmType1Config) *entity.ComputePriceResult{
	fmt.Println("param.Year=",param.Year)
	fmt.Println("config.FeePerYear=",config.FeePerYear)
	fmt.Println("config.TimeNodes=",config.TimeNodes)
	fmt.Println("config.Discounts=",config.Discounts)

	ret := &entity.ComputePriceResult{
		OldPrice: float32(param.Year) * config.FeePerYear,
		Discount: config.Discounts[0],
		NewPrice: float32(param.Year) * config.FeePerYear}

	if(0 == len(config.TimeNodes)){
		return ret
	}

	rangIndex := 0
	nodeCount := len(config.TimeNodes)
	for i := 0 ; i <  nodeCount; i++ {
		if(param.Year >= config.TimeNodes[i]){
			if(i >= nodeCount-1){
				rangIndex = nodeCount - 1
			}else{
				continue;
			}
		}else{
			rangIndex = i-1;
			break;
		}
	}
	fmt.Println("rangIndex=",rangIndex)
	
	ret.Discount = config.Discounts[rangIndex]
	ret.NewPrice = ret.OldPrice * ret.Discount

	fmt.Println("ret=",ret)

	return ret
}