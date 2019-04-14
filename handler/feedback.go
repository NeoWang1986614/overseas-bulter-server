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

const (
	feedback_read 		= 0
	feedback_unread 	= 1
	feedback_unvalid 	= 2
)

func FeedbackSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("feedback search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postFeedbackSearchHandler(w, r)
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

func FeedbackHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("feedback handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postFeedbackHandler(w, r)
		break;
	case "PUT":
		// putHouseHandler(w, r)
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func postFeedbackHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Feedback{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddFeedback(
		requestBody.OrderId,
		requestBody.AuthorId,
		requestBody.Content)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postFeedbackSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.FeedbackQuery{IsRead: feedback_unvalid}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)
	
	arr := make([]storage.DbFeedback, 0)
	if(requestBody.IsRead == feedback_unvalid){
		arr = storage.QueryFeedbackByOrderId(requestBody.Length, requestBody.Offset, requestBody.OrderId);
	}else{
		arr = storage.QueryFeedbackByOrderIdIsRead(requestBody.Length, requestBody.Offset, requestBody.OrderId, requestBody.IsRead);
	}
	
	entities := make([]entity.Feedback, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Feedback{
			arr[i].Uid,
			arr[i].OrderId,
			arr[i].AuthorId,
			arr[i].Content,
			arr[i].IsRead,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))

}
