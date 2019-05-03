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

func RepairRecordHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getRepairRecordHandler(w, r)
		break;
	case "POST":
		postRepairRecordHandler(w, r)
		break;
	case "PUT":
		putRepairRecordHandler(w, r)
		break;
	case "DELETE":
		deleteRepairRecordHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func RepairRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postRepairRecordSearchHandler(w, r)
		break;
	case "PUT":
		putRepairRecordHandler(w, r)
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func getRepairRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbRecord := storage.QueryRepairRecord(uid[0])
	entity := entity.ConvertToRepairRecordEntity(dbRecord)
	rsp, err := json.Marshal(entity)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postRepairRecordHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.RepairRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddRepairRecord(
		requestBody.OrderId,
		requestBody.ReportTime,
		requestBody.RepairTime,
		requestBody.CompleteTime,
		requestBody.Comment,
		requestBody.Cost,
		requestBody.Status,
		requestBody.RelatedImage)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func putRepairRecordHandler(w http.ResponseWriter, r *http.Request)  {

	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("uid = ", uid)

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.RepairRecord{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateRepairRecord(
		uid[0],
		requestBody.OrderId,
		requestBody.ReportTime,
		requestBody.RepairTime,
		requestBody.CompleteTime,
		requestBody.Comment,
		requestBody.Cost,
		requestBody.Status,
		requestBody.RelatedImage)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteRepairRecordHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("delete uid : ", uid);
	storage.DeleteRepairRecordByUid(uid[0])
	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postRepairRecordSearchHandler(w http.ResponseWriter, r *http.Request)  {
	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.RepairRecordQueryByOrderId{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	total, arr := storage.QueryRepairRecordByOrderId(requestBody.OrderId, requestBody.Offset, requestBody.Count)

	entities := make([]entity.RepairRecord, 0)
	if(0 != total){
		for i := 0 ; i < len(arr) ; i ++ {
			var enti = entity.ConvertToRepairRecordEntity(&arr[i])
			entities = append(entities, *enti)
		}
	}
	
	qRet := &entity.RepairRecordQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}