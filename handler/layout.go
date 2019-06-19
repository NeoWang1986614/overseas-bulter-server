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

func LayoutHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getLayoutHandler(w, r)
		break;
	case "POST":
		postLayoutHandler(w, r)
		break;
	case "PUT":
		putLayoutHandler(w, r)
		break;
	case "DELETE":
		deleteLayoutHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func LayoutSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postLayoutSearchHandler(w, r)
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

func postLayoutHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Layout{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	err = storage.AddLayout(
		requestBody.Value,
		requestBody.Title,
		requestBody.Location,
		requestBody.Content,
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

func getLayoutHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbResult := storage.QueryLayout(uid[0])

	enti := entity.ConvertToLayoutEntity(dbResult)

	rsp, err := json.Marshal(enti)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putLayoutHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Layout{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateLayout(
		requestBody.Uid,
		requestBody.Value,
		requestBody.Title,
		requestBody.Location,
		requestBody.Content,
		requestBody.Meta)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteLayoutHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	storage.DeleteLayout(uid[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postLayoutSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("postLayoutSearchHandler");

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.LayoutRangeQuery{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	arr := storage.QueryLayoutByRange(requestBody.Count, requestBody.Offset)

	entities := make([]entity.Layout, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		enti := entity.ConvertToLayoutEntity(&arr[i])
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}