package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	// "encoding/json"
	wxpay "github.com/objcoding/wxpay"
	storage "overseas-bulter-server/storage"
	// entity "overseas-bulter-server/entity"
	// Error "overseas-bulter-server/error"
	// entity "overseas-bulter-server/entity"
)

func WxPayNotifyHandler(w http.ResponseWriter, r *http.Request)  {
	
	fmt.Println("wx pay handler")
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	con,_:=ioutil.ReadAll(r.Body)
	// fmt.Println(string(con))
	if("POST" == r.Method){
		params := wxpay.XmlToMap(string(con))
		fmt.Println("notify content: ",string(con))
		if("SUCCESS" == params["result_code"] && "SUCCESS"== params["return_code"]){
			orderId := params["out_trade_no"]
			storage.UpdateOrderStatusByUid("paid", orderId)
			 CORSHandle(w)
			 noti := &wxpay.Notifies{}
			io.WriteString(w, noti.OK())
		}
	}
}
