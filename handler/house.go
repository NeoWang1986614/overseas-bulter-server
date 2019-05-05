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
	dbHouse := storage.QueryHouse(id[0])
	house := entity.ConvertToHouseEntity(dbHouse)
	rsp, err := json.Marshal(house)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	io.WriteString(w, string(rsp))
}

func postHouseHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("post house handler++");
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.House{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	uid := storage.AddHouse(
		requestBody.Name,
		requestBody.Lat,
		requestBody.Lng,
		requestBody.AdLevel1,
		requestBody.AdLevel2,
		requestBody.AdLevel3,
		requestBody.Locality,
		requestBody.Nation,
		requestBody.StreetName,
		requestBody.StreetNum,
		requestBody.BuildingNum,
		requestBody.RoomNum,
		requestBody.Layout,
		requestBody.Area,
		requestBody.OwnerId)

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
		requestBody.Lat,
		requestBody.Lng,
		requestBody.AdLevel1,
		requestBody.AdLevel2,
		requestBody.AdLevel3,
		requestBody.Locality,
		requestBody.Nation,
		requestBody.StreetName,
		requestBody.StreetNum,
		requestBody.BuildingNum,
		requestBody.RoomNum,
		requestBody.Layout,
		requestBody.Area,
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
		var enti = entity.ConvertToHouseEntity(&arr[i])
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
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
