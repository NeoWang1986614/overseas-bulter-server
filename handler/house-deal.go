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

func HouseDealHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getHouseDealHandler(w, r)
		break;
	case "POST":
		postHouseDealHandler(w, r)
		break;
	case "PUT":
		putHouseDealHandler(w, r)
		break;
	case "DELETE":
		deleteHouseDealHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func HouseDealSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postHouseDealSearchHandler(w, r)
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

func postHouseDealHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.HouseDeal{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	err = storage.AddHouseDeal(
		requestBody.DealType,
		requestBody.Source,
		requestBody.HouseId,
		requestBody.Decoration,
		requestBody.Cost,
		requestBody.Linkman,
		requestBody.ContactNum,
		requestBody.Mail,
		requestBody.Weixin,
		requestBody.Image,
		requestBody.Note,
		requestBody.Creator,
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

func getHouseDealHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbResult := storage.QueryHouseDeal(uid[0])

	enti := entity.ConvertToHouseDealEntity(dbResult)

	rsp, err := json.Marshal(enti)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putHouseDealHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.HouseDeal{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateHouseDeal(
		requestBody.Uid,
		requestBody.DealType,
		requestBody.Source,
		requestBody.HouseId,
		requestBody.Decoration,
		requestBody.Cost,
		requestBody.Linkman,
		requestBody.ContactNum,
		requestBody.Mail,
		requestBody.Weixin,
		requestBody.Image,
		requestBody.Note,
		requestBody.Creator,
		requestBody.Meta)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteHouseDealHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	storage.DeleteHouseDeal(uid[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postHouseDealSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("postHouseDealSearchHandler");

	qType, ok := r.URL.Query()["queryType"]
	if(!ok) {
		panic("no qType in url param")
	}
	fmt.Println("qType = ", qType)

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	var total uint = 0
	arr := make([]storage.DbHouseDeal, 0)
	if("range" == qType[0]){
		requestBody := &entity.HouseDealRangeQuery{}
		err := json.Unmarshal(con, requestBody)
		Error.CheckErr(err)
		fmt.Print(requestBody)
		arr = storage.QueryHouseDealByRange(requestBody.Count, requestBody.Offset)

	}else if("deal_type"==qType[0]){
		requestBody := &entity.HouseDealQueryByDealType{}
		err := json.Unmarshal(con, requestBody)
		Error.CheckErr(err)
		fmt.Print(requestBody)
		total, arr = storage.QueryHouseDealByDealType(requestBody.DealType ,requestBody.Count, requestBody.Offset)
	}
	
	entities := make([]entity.HouseDeal, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		enti := entity.ConvertToHouseDealEntity(&arr[i])
		entities = append(entities, *enti)
	}

	ret := &entity.HouseDealQueryResult{
		Total: total,
		Entities: entities}


	rsp, err := json.Marshal(ret)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}