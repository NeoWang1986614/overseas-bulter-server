package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	// // wx "overseas-bulter-server/wx"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
	storage "overseas-bulter-server/storage"
	// entity "overseas-bulter-server/entity"
)
func RentHouseSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postRentHouseSearchHandler(w, r)
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

func makeRentHouseEntity(entity *entity.RentHouse,  inspectRecords []storage.DbInspectRecord, house *storage.DbHouse, order *storage.DbOrder){
	fmt.Println("makeRentHouseEntity");

	if(0!=len(inspectRecords)){
		entity.HouseImage = inspectRecords[0].Area
	}
	
	entity.OrderId =  order.Uid
	entity.HouseName = house.Name
	entity.HouseLayout = house.Layout
	entity.HouseAdLevel2 = house.AdLevel2
	entity.HouseAdLevel3 = house.AdLevel3
	entity.OrderMeta = order.Meta

}

func postRentHouseSearchHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	requestBody := &entity.QuerySearchRentHouse{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)


	houseRentOrderTypeArr := []string{"house-rent"}
	total, orderArr := storage.QueryOrderByOrderTypeGroupStatus(houseRentOrderTypeArr, "accepted", requestBody.Offset, requestBody.Count)

	entities := make([]entity.RentHouse, 0)
	if(0 != len(orderArr)){
		for i := 0 ; i < len(orderArr) ; i ++ {
			entity := &entity.RentHouse{};
			_, arr := storage.QueryInspectRecordByOrderId(orderArr[i].Uid, 0, 10000)
			h := storage.QueryHouse(orderArr[i].HouseId)
			makeRentHouseEntity(entity, arr, h, &orderArr[i])
			entities = append(entities, *entity)
		}	
	}
	
	ret := &entity.RentHouseQueryResult{
		Total: total,
	Entities: entities}

	rsp, err := json.Marshal(ret)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}