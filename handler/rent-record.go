package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	storage "overseas-bulter-server/storage"
	Error "overseas-bulter-server/error"
	entity "overseas-bulter-server/entity"
)

func RentRecordHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("public account material handler")
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getRentRecordHandler(w, r)
		break;
	case "POST":
		postRentRecordHandler(w, r)
		break;
	case "PUT":
		putRentRecordHandler(w, r)
		break;
	case "DELETE":
		deleteRentRecordHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func RentRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("public account material search handler")
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postRentRecordSearchHandler(w, r)
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

func getRentRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbRecord := storage.QueryRentRecord(uid[0])
	entity := entity.ConvertToRentRecordEntity(dbRecord)
	rsp, err := json.Marshal(entity)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postRentRecordHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.RentRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddRentRecord(
		requestBody.OrderId,
		requestBody.Income,
		requestBody.Outgoings,
		requestBody.Balance,
		requestBody.Comment,
		requestBody.TimeRange,
		requestBody.AccountingDate)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func putRentRecordHandler(w http.ResponseWriter, r *http.Request)  {

	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.RentRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateRentRecord(
		uid[0],
		requestBody.OrderId,
		requestBody.Income,
		requestBody.Outgoings,
		requestBody.Balance,
		requestBody.Comment,
		requestBody.TimeRange,
		requestBody.AccountingDate)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteRentRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("delete uid : ", uid);
	storage.DeleteRentRecordByUid(uid[0])
	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postRentRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.RentRecordQueryByOrderId{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	total, arr := storage.QueryRentRecordByOrderId(requestBody.OrderId, requestBody.Offset, requestBody.Count)

	entities := make([]entity.RentRecord, 0)
	if(0 != total){
		for i := 0 ; i < len(arr) ; i ++ {
			var enti = entity.ConvertToRentRecordEntity(&arr[i])
			entities = append(entities, *enti)
		}
	}
	
	qRet := &entity.RentRecordQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}