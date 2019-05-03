package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	wx "overseas-bulter-server/wx"
	// storage "overseas-bulter-server/storage"
	Error "overseas-bulter-server/error"
	entity "overseas-bulter-server/entity"
)

func PublicAccountMaterialHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("public account material handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getPublicAccountMaterailHandler(w, r)
		break;
	case "POST":
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

func PublicAccountMaterialSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("public account material search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postPublicAccountMaterailSearchHandler(w, r)
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

func getPublicAccountMaterailHandler(w http.ResponseWriter, r *http.Request)  {
	mId, ok := r.URL.Query()["media_id"]
	if(!ok) {
		panic("no mId exist in url param")
	}
	fmt.Println("id = ", mId)
	materialDetail := wx.GetPublicAccountMaterailDetail(mId[0])
	fmt.Println("get material detail: ", materialDetail)
	rsp, err := json.Marshal(materialDetail)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	io.WriteString(w, string(rsp))
}

func postPublicAccountMaterailSearchHandler(w http.ResponseWriter, r *http.Request)  {
	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.PublicAccountMaterialQuery{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	material := wx.GetPublicAccountMaterail(requestBody.Type, requestBody.Offset, requestBody.Count)
	// fmt.Println("material = ", material);

	rsp, err := json.Marshal(material)
	Error.CheckErr(err)
	// fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}