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
	// entity "overseas-bulter-server/entity"
)

func CaseSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("case search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postCaseSearchHandler(w, r)
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

func CaseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("case handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getCaseHandler(w, r)
		break;
	case "POST":
		postCaseHandler(w, r)
		break;
	case "PUT":
		putCaseHandler(w, r)
		break;
	case "DELETE":
		deleteCaseHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func getCaseHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("id = ", id)
	enti := storage.QueryCase(id[0])
	caseEnti := &entity.Case{
		Uid: enti.Uid,
		Title: enti.Title,
		ImageUrl: enti.ImageUrl,
		Content: enti.Content,
		Price: enti.Price,
		Level: enti.Level}

	rsp, err := json.Marshal(caseEnti)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postCaseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Case{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	uid := storage.AddCase(
		requestBody.Title,
		requestBody.ImageUrl,
		requestBody.Content,
		requestBody.Price,
		requestBody.Level)
		
	requestBody.Uid = uid;
	rsp, err := json.Marshal(requestBody)
	fmt.Println(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putCaseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Case{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateCase(
		requestBody.Uid,
		requestBody.Title,
		requestBody.ImageUrl,
		requestBody.Content,
		requestBody.Price,
		requestBody.Level)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteCaseHandler(w http.ResponseWriter, r *http.Request)  {
	id, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("delete index : ", id);
	storage.DeleteCaseByUid(id[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postCaseSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.CaseSearchByLevel{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	arr := storage.QueryCasesByLevel(requestBody.Length, requestBody.Offset, requestBody.Level);

	entities := make([]entity.Case, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Case{
			arr[i].Uid,
			arr[i].Title,
			arr[i].ImageUrl,
			arr[i].Content,
			arr[i].Price,
			arr[i].Level}
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))

}
