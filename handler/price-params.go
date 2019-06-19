package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	// wx "overseas-bulter-server/wx"
	storage "overseas-bulter-server/storage"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
)

func PriceParamsHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getPriceParamsHandler(w, r)
		break;
	case "POST":
		postPriceParamsHandler(w, r)
		break;
	case "PUT":
		putPriceParamsHandler(w, r)
		break;
	case "DELETE":
		deletePriceParamsHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func PriceParamsSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postPriceParamsSearchHandler(w, r)
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

func postPriceParamsHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.PriceParams{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	err = storage.AddPriceParams(
		requestBody.ServiceId,
		requestBody.LayoutId,
		requestBody.AlgorithmType,
		requestBody.Params,
		requestBody.Meta)
	
	ret := ""
	if(nil != err){
		ret = GetErrJsonString(1, "错误: 添加出错!")
	}else{
		ret = GetSuccessJsonString()
	}

	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func getPriceParamsHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbResult := storage.QueryPriceParams(uid[0])

	enti := entity.ConvertToPriceParamsEntity(dbResult)

	rsp, err := json.Marshal(enti)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putPriceParamsHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.PriceParams{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdatePriceParams(
		requestBody.Uid,
		requestBody.ServiceId,
		requestBody.LayoutId,
		requestBody.AlgorithmType,
		requestBody.Params,
		requestBody.Meta)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deletePriceParamsHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	storage.DeletePriceParams(uid[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postPriceParamsSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("postPriceParamsSearchHandler");

	queryType, ok := r.URL.Query()["queryType"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("queryType = ", queryType)
	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	arr := make([]storage.DbPriceParams, 0)
	if("range" == queryType[0]){
		requestBody := &entity.PriceParamsRangeQuery{}
		err := json.Unmarshal(con, requestBody)
		Error.CheckErr(err)
		fmt.Print(requestBody)
		arr = storage.QueryPriceParamsByRange(requestBody.Count, requestBody.Offset)
	}else if("service-layout" == queryType[0]){
		requestBody := &entity.PriceParamsServiceIdLayoutIdQuery{}
		err := json.Unmarshal(con, requestBody)
		Error.CheckErr(err)
		fmt.Print(requestBody)
		arr = storage.QueryPriceParamsByServiceIdLayoutId(requestBody.ServiceId, requestBody.LayoutId)
	}

	entities := make([]entity.PriceParams, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		enti := entity.ConvertToPriceParamsEntity(&arr[i])
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}