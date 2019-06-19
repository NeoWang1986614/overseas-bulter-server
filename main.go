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
	// wx "overseas-bulter-server/wx"
)

func main() {

	//debug
	// uuid
	uuid := uuid.GenerateNextUuid()
	fmt.Println(uuid)

	//config
	config.Init();
	//storage
	storage.Init()
	// wx.CreateAccount()
	
	//http
	http.HandleFunc(config.GenerateIntegratedUri("/login"), handler.LoginHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/user"), handler.UserHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/house"), handler.HouseHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/house/search"), handler.HouseSearchHandler)
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
	http.HandleFunc(config.GenerateIntegratedUri("/billing-record"), handler.BillingRecordHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/billing-record/search"), handler.BillingRecordSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/inspect-record"), handler.InspectRecordHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/inspect-record/search"), handler.InspectRecordSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/repair-record"), handler.RepairRecordHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/repair-record/search"), handler.RepairRecordSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/price"), handler.PriceHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/prepayment"), handler.PrepaymentHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/wxpay/notify"), handler.WxPayNotifyHandler)

	http.HandleFunc(config.GenerateIntegratedUri("/service-primary"), handler.ServciePrimaryHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/service-primary/search"), handler.ServciePrimarySearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/layout"), handler.LayoutHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/layout/search"), handler.LayoutSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/price-params"), handler.PriceParamsHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/price-params/search"), handler.PriceParamsSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/rent-house/search"), handler.RentHouseSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/carousel-figure"), handler.CarouselFigureHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/carousel-figure/search"), handler.CarouselFigureSearchHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/house-deal"), handler.HouseDealHandler)
	http.HandleFunc(config.GenerateIntegratedUri("/house-deal/search"), handler.HouseDealSearchHandler)


	http.HandleFunc(config.GenerateIntegratedUri("/file"), handler.FileHandler)

	fmt.Println("start https server, listen on ", config.ServerListentIp, "...")
	// err := http.ListenAndServe(config.ServerListentIp,nil)
	// err := http.ListenAndServeTLS(config.ServerListentIp, "1_cqmygysdss.com_bundle.crt","2_cqmygysdss.com.key", nil)
	err := http.ListenAndServeTLS(config.ServerListentIp, "2076716_bulter.mroom.com.cn.crt","2076716_bulter.mroom.com.cn.key", nil)
	if err != nil {
        log.Fatal("start https server err: ", err)
	}
	
	fmt.Println("Exit !")
}