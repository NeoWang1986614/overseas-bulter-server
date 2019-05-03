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
	http.HandleFunc(config.GenerateIntegratedUri("/employee"), handler.EmployeeHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/employee/search"), handler.EmployeeSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/employee/check"), handler.EmployeeCheckHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/image"), handler.ImageHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/share"), handler.ShareHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/public-account/material/search"), handler.PublicAccountMaterialSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/public-account/material"), handler.PublicAccountMaterialHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/rent-record"), handler.RentRecordHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/rent-record/search"), handler.RentRecordSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/inspect-record"), handler.InspectRecordHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/inspect-record/search"), handler.InspectRecordSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/repair-record"), handler.RepairRecordHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/repair-record/search"), handler.RepairRecordSearchHandler)

	fmt.Println("start http server...")
	// err := http.ListenAndServe(config.ServerListentIp,nil)
	// err := http.ListenAndServeTLS(config.ServerListentIp, "1_cqmygysdss.com_bundle.crt","2_cqmygysdss.com.key", nil)
	err := http.ListenAndServeTLS(config.ServerListentIp, "2076716_bulter.mroom.com.cn.crt","2076716_bulter.mroom.com.cn.key", nil)
	if err != nil {
        log.Fatal("start http server err: ", err)
	}
	
	fmt.Println("Exit !")
}