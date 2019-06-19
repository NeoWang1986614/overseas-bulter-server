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

func ServciePrimaryHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getServicePrimaryHandler(w, r)
		break;
	case "POST":
		postServicePrimaryHandler(w, r)
		break;
	case "PUT":
		putServicePrimaryHandler(w, r)
		break;
	case "DELETE":
		deleteServicePrimaryHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func ServciePrimarySearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postServicePrimarySearchHandler(w, r)
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

func postServicePrimaryHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.ServicePrimary{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	err = storage.AddServicePrimay(
		requestBody.Value,
		requestBody.Title,
		requestBody.IconUrl,
		requestBody.Location,
		requestBody.Content,
		requestBody.BasePrice,
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

func getServicePrimaryHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbResult := storage.QueryServicePrimary(uid[0])

	enti := entity.ConvertToServicePrimaryEntity(dbResult)

	rsp, err := json.Marshal(enti)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putServicePrimaryHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.ServicePrimary{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateServicePrimary(
		requestBody.Uid,
		requestBody.Value,
		requestBody.Title,
		requestBody.IconUrl,
		requestBody.Location,
		requestBody.Content,
		requestBody.BasePrice,
		requestBody.Meta)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteServicePrimaryHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	storage.DeleteServicePrimary(uid[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postServicePrimarySearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("postServicePrimarySearchHandler");

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.ServicePrimaryRangeQuery{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	arr := storage.QueryServicePrimaryByRange(requestBody.Count, requestBody.Offset)

	entities := make([]entity.ServicePrimary, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		enti := entity.ConvertToServicePrimaryEntity(&arr[i])
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}