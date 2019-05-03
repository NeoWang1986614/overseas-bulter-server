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

func InspectRecordHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getInspectRecordHandler(w, r)
		break;
	case "POST":
		postInspectRecordHandler(w, r)
		break;
	case "PUT":
		putInspectRecordHandler(w, r)
		break;
	case "DELETE":
		deleteInspectRecordHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func InspectRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postInspectRecordSearchHandler(w, r)
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

func getInspectRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbRecord := storage.QueryInspectRecord(uid[0])
	entity := entity.ConvertToInspectRecordEntity(dbRecord)
	rsp, err := json.Marshal(entity)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postInspectRecordHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.InspectRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddInspectRecord(
		requestBody.OrderId,
		requestBody.InspectDate,
		requestBody.Inspector,
		requestBody.Comment,
		requestBody.Config,
		requestBody.Area)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func putInspectRecordHandler(w http.ResponseWriter, r *http.Request)  {

	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.InspectRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateInspectRecord(
		uid[0],
		requestBody.OrderId,
		requestBody.InspectDate,
		requestBody.Inspector,
		requestBody.Comment,
		requestBody.Config,
		requestBody.Area)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteInspectRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("delete uid : ", uid);
	storage.DeleteInspectRecordByUid(uid[0])
	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postInspectRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.InspectRecordQueryByOrderId{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	total, arr := storage.QueryInspectRecordByOrderId(requestBody.OrderId, requestBody.Offset, requestBody.Count)

	entities := make([]entity.InspectRecord, 0)
	if(0 != total){
		for i := 0 ; i < len(arr) ; i ++ {
			var enti = entity.ConvertToInspectRecordEntity(&arr[i])
			entities = append(entities, *enti)
		}
	}
	
	qRet := &entity.InspectRecordQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}