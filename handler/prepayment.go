package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	wx "overseas-bulter-server/wx"
	// storage "overseas-bulter-server/storage"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"

)

func PrepaymentHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("pre payment handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		PostPrepaymentHandler(w, r)
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

func PostPrepaymentHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Prepayment{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	// dbOrder := storage.QueryOrder(requestBody.OrderId)

	prepaymentRet, err := wx.Prepayment(requestBody.UserId, requestBody.OrderId, 1/*int64(dbOrder.Price) * 100*/)

	if(nil != err){
		errEntity := entity.GetErr(1007, "支付错误!")
		rsp, err := json.Marshal(errEntity)
		Error.CheckErr(err)
		fmt.Print(string(rsp))
		CORSHandle(w)
		io.WriteString(w, string(rsp))
		Error.CheckErr(err)
	}

	rsp, err := json.Marshal(prepaymentRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}