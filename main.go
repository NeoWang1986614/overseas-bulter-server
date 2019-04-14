package main

import (
	"fmt"
	"net/http"
	"log"
	// "io"
	// "overseas-bulter-server/entity"
	// "encoding/json"
	config "overseas-bulter-server/config"
	uuid "overseas-bulter-server/uuid"
	storage "overseas-bulter-server/storage"
	handler "overseas-bulter-server/handler"
	// login "overseas-bulter-server/login"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler");
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login handler");
	// io.WriteString(w, string(getMockData()))
}

// func getMockData() []byte{
// 	test := new(entity.Customer)
// 	test.Id = "34234"
// 	test.Age = 32
// 	test.Name = "你大爷"
// 	test.Sex = "female"

// 	test1 := new(entity.Customer)
// 	test1.Id = "34234"
// 	test1.Age = 32
// 	test1.Name = "你大爷"
// 	test1.Sex = "female"

// 	result := []entity.Asset{
// 		entity.Asset {
// 			"3213213",
// 			"我的资产1",
// 			"xxx小区xx栋xx路xx小区xx层xx室",
// 			"321323123",
// 			1,
// 			"893128739",
// 		},
// 		entity.Asset {
// 			"3213214",
// 			"我的资产2",
// 			"xxx小区xx栋xx路xx小区xx层xx室",
// 			"321323123",
// 			1,
// 			"893128739",
// 		},
// 		entity.Asset {
// 			"3213215",
// 			"我的资产3",
// 			"xxx小区xx栋xx路xx小区xx层xx室",
// 			"321323123",
// 			1,
// 			"893128739",
// 		},
// 	}

// 	fmt.Println(test);

// 	b, e := json.Marshal(result)
// 	if e == nil {
// 		fmt.Println(string(b))
// 	} else {
// 		fmt.Println(e)
// 	}
// 	return b;
// }

func main() {

	//debug
	// uuid
	uuid := uuid.GenerateNextUuid()
	fmt.Println(uuid)

	//storage
	storage.Init()

	//http
	http.HandleFunc(config.GenerateIntegratedUri("/login"), handler.LoginHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/user"), handler.UserHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/house"), handler.HouseHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/house/search"), handler.HouseSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/case"), handler.CaseHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/case/search"), handler.CaseSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/order"), handler.OrderHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/order/search"), handler.OrderSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/order/search/advanced"), handler.OrderSearchAdvancedHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/feedback"), handler.FeedbackHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/feedback/search"), handler.FeedbackSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/service"), handler.ServcieHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/service/search"), handler.ServcieSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/image"), handler.ImageHandler)

	fmt.Println("start http server...")
	err := http.ListenAndServe(config.ServerListentIp,nil)
	// err := http.ListenAndServeTLS(config.ServerListentIp, "1_cqmygysdss.com_bundle.crt","2_cqmygysdss.com.key", nil)
	if err != nil {
        log.Fatal("start http server err: ", err)
	}
	
	fmt.Println("Exit !")
}