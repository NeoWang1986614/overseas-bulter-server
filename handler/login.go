package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	wx "overseas-bulter-server/wx"
	storage "overseas-bulter-server/storage"
	Error "overseas-bulter-server/error"
	entity "overseas-bulter-server/entity"
)

type LoginRequestBody struct {
	Code string
}

func LoginHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("login handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postLoginHandler(w, r)
		break;
	case "PUT":
		break;
	case "DELETE":
		break;
	}	
}

func postLoginHandler(w http.ResponseWriter, r *http.Request) {

	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &LoginRequestBody{}
	
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	sessionObject := wx.Code2Session(requestBody.Code)
	fmt.Println(sessionObject.OpenId, sessionObject.Session_key)
	storage.UpdateUserByWxOpenId(sessionObject.OpenId, sessionObject.Session_key)
	user := storage.QueryUserByWxOpenId(sessionObject.OpenId)
	
	userRsp := &entity.User{
		Uid: user.Uid,
		IdCardNumber: user.IdCardNumber,
		PhoneNumber: user.PhoneNumber}
	

	rsp, _ := json.Marshal(userRsp)
	
	fmt.Println(string(rsp))

	io.WriteString(w, string(rsp))
}