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

func ServcieHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("service handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getServiceHandler(w, r)
		break;
	case "POST":
		postServiceHandler(w, r)
		break;
	case "PUT":
		putServiceHandler(w, r)
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func ServcieSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("service search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postServiceSearchHandler(w, r)
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

func postServiceSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.ServiceQuery{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	arr := storage.QueryServicesByType(requestBody.Type)

	entities := make([]entity.Service, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Service{
			arr[i].Uid,
			arr[i].Type,
			arr[i].Layout,
			arr[i].Content,
			arr[i].Price,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postServiceHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Service{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddService(
		requestBody.Type,
		requestBody.Layout,
		requestBody.Content,
		requestBody.Price)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func getServiceHandler(w http.ResponseWriter, r *http.Request)  {
	id, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("id = ", id)
	enti := storage.QueryService(id[0])

	service := &entity.Service{
		Uid: enti.Uid,
		Type: enti.Type,
		Layout: enti.Layout,
		Content: enti.Content,
		Price: enti.Price,
		CreateTime: enti.CreateTime}
	rsp, err := json.Marshal(service)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	io.WriteString(w, string(rsp))
}

func putServiceHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Service{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateService(
		requestBody.Uid,
		requestBody.Type,
		requestBody.Layout,
		requestBody.Content,
		requestBody.Price)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}
