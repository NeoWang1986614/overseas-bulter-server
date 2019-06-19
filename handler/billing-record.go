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

func BillingRecordHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("billing record handler")
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getBillingRecordHandler(w, r)
		break;
	case "POST":
		postBillingRecordHandler(w, r)
		break;
	case "PUT":
		putBillingRecordHandler(w, r)
		break;
	case "DELETE":
		deleteBillingRecordHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func BillingRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("public account material search handler")
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postBillingRecordSearchHandler(w, r)
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

func getBillingRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbRecord := storage.QueryBillingRecord(uid[0])
	entity := entity.ConvertToBillingRecordEntity(dbRecord)
	rsp, err := json.Marshal(entity)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postBillingRecordHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.BillingRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddBillingRecord(
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

func putBillingRecordHandler(w http.ResponseWriter, r *http.Request)  {

	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.BillingRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateBillingRecord(
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

func deleteBillingRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("delete uid : ", uid);
	storage.DeleteBillingRecordByUid(uid[0])
	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postBillingRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.BillingRecordQueryByOrderId{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	total, arr := storage.QueryBillingRecordByOrderId(requestBody.OrderId, requestBody.Offset, requestBody.Count)

	entities := make([]entity.BillingRecord, 0)
	if(0 != total){
		for i := 0 ; i < len(arr) ; i ++ {
			var enti = entity.ConvertToBillingRecordEntity(&arr[i])
			entities = append(entities, *enti)
		}
	}
	
	qRet := &entity.BillingRecordQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}