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

func HouseSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("house search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postHouseSearchHandler(w, r)
		break;
	case "PUT":
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		fmt.Println(r.Header.Get("Content-Type"))
		CORSHandle(w)
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// w.Header().Set("Content-Type", "application/json;charset=utf-8")
		break;
	}	
}

func HouseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("house handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getHouseHandler(w, r)
		break;
	case "POST":
		postHouseHandler(w, r)
		break;
	case "PUT":
		putHouseHandler(w, r)
		break;
	case "DELETE":
		deleteHouseHandler(w, r)
		break;
	}	
}

func getHouseHandler(w http.ResponseWriter, r *http.Request)  {
	id, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("id = ", id)
	enti := storage.QueryHouse(id[0])
	order := &entity.House{
		Uid: enti.Uid,
		Name: enti.Name,
		Country: enti.Country,
		Province: enti.Province,
		City: enti.City,
		Address: enti.Address,
		Layout: enti.Layout,
		OwnerId: enti.OwnerId}
	rsp, err := json.Marshal(order)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	io.WriteString(w, string(rsp))
}

func postHouseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.House{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	uid := storage.AddHouse(
		requestBody.Name,
		requestBody.Country,
		requestBody.Province,
		requestBody.City,
		requestBody.Address,
		requestBody.Layout,
		requestBody.OwnerId)

		// AddUserResult
	ret := &entity.AddHouseResult{Uid: uid}

	rsp, err := json.Marshal(ret)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	io.WriteString(w, string(rsp))
}

func putHouseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.House{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateHouse(
		requestBody.Uid,
		requestBody.Name,
		requestBody.Country,
		requestBody.Province,
		requestBody.City,
		requestBody.Address,
		requestBody.Layout,
		requestBody.OwnerId)
	
	ret := &entity.AddHouseResult{Uid: requestBody.Uid}

	rsp, err := json.Marshal(ret)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	io.WriteString(w, string(rsp))
	
}

func postHouseSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.HouseSearch{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	entities := make([]entity.House, 0)
	arr := storage.QueryHouses(requestBody.OwnerId, requestBody.Length, requestBody.Offset);

	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.House{
			arr[i].Uid,
			arr[i].Name,
			arr[i].Country,
			arr[i].Province,
			arr[i].City,
			arr[i].Address,
			arr[i].Layout,
			arr[i].OwnerId}
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json;charset=utf-8")
	io.WriteString(w, string(rsp))

}

func deleteHouseHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("delete index : ", uid);
	storage.DeleteHouseByUid(uid[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	io.WriteString(w, ret)
}
