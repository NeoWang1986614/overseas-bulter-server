package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	// wx "overseas-bulter-server/wx"
	wx "overseas-bulter-server/wx"
	storage "overseas-bulter-server/storage"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
	// entity "overseas-bulter-server/entity"
)

func ShareHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("share handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		// getUserHandler(w, r)
		break;
	case "POST":
		postShareHandler(w, r)
		break;
	case "PUT":
		// putUserHandler(w, r)
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func postShareHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.AddShareRequest{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	user := storage.QueryUserByUid(requestBody.FromUid)
	
	fmt.Println("from user = ", user);

	shareInfo := wx.DecodeShareInfo(user.WxSessionKey, requestBody.EncryptedData, requestBody.Iv)
	
	fmt.Println("shareInfo = ", shareInfo);

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

// func UserSearchHandler(w http.ResponseWriter, r *http.Request)  {
// 	fmt.Println("user search handler")
// 	fmt.Println(r.Method);
// 	switch(r.Method){
// 	case "GET":
// 		break;
// 	case "POST":
// 		postSearchHandler(w, r)
// 		break;
// 	case "PUT":
// 		break;
// 	case "DELETE":
// 		break;
// 	case "OPTIONS":
// 		CORSHandle(w)
// 		break;
// 	}	
// }


// func getUserHandler(w http.ResponseWriter, r *http.Request)  {
// 	uid, ok := r.URL.Query()["uid"]
// 	if(!ok) {
// 		panic("no uid exist in url param")
// 	}

// 	enti := storage.QueryUserByUid(uid[0])
// 	user := &entity.User{
// 		Uid: enti.Uid,
// 		Name: enti.Name,
// 		PhoneNumber: enti.PhoneNumber,
// 		IdCardNumber: enti.IdCardNumber}

// 	rsp, err := json.Marshal(user)
// 	Error.CheckErr(err)
// 	fmt.Print(string(rsp))
// 	CORSHandle(w)
// 	io.WriteString(w, string(rsp))
// }

// func putUserHandler(w http.ResponseWriter, r *http.Request)  {
// 	fmt.Println(r.Body);

// 	con,_:=ioutil.ReadAll(r.Body)
// 	fmt.Println(string(con))

// 	requestBody := &entity.User{}
// 	err := json.Unmarshal(con, requestBody)
// 	Error.CheckErr(err)
// 	fmt.Print(requestBody)

// 	storage.UpdateUserByUid(
// 		requestBody.Uid,
// 		requestBody.Name,
// 		requestBody.PhoneNumber,
// 		requestBody.IdCardNumber)

// 	ret := GetSuccessJsonString()
// 	fmt.Println(ret)
// 	CORSHandle(w)
// 	io.WriteString(w, ret)
// }

// func postSearchHandler(w http.ResponseWriter, r *http.Request) {

// }
