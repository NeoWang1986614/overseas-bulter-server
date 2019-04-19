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
	// entity "overseas-bulter-server/entity"
)

const(
	orderQueryTypeByIdCardNumber = "id_card_number"
	orderQueryTypeByPhoneNumber = "phone_number"
	orderQueryTypeByRealName = "real_name"
	orderQueryTypeBeforeTime = "before_time"
	orderQueryTypeAfterTime = "after_time"
	orderQueryTypeRangeTime = "range_time"
	orderQueryTypeByStatusGroup = "status_group"
	orderQueryTypeByAddress = "address"
	orderQueryTypeByLayoutGroup = "layout_group"
	orderQueryTypeBelowPrice = "below_price"
	orderQueryTypeAbovePrice = "above_price"
	orderQueryTypeRangePrice = "range_price"
	orderQueryTypeByOrderTypeGroup = "order_type_group"
	orderQueryTypeAll = "all"
)

func OrderSearchAdvancedHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("order search advanced handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postOrderSearchAdvancedHandler(w, r)
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

func OrderSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("order search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postOrderSearchHandler(w, r)
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

func OrderHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("order handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getOrderHandler(w, r)
		break;
	case "POST":
		postOrderHandler(w, r)
		break;
	case "PUT":
		putOrderHandler(w, r)
		break;
	case "DELETE":
		deleteOrderHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func getOrderHandler(w http.ResponseWriter, r *http.Request)  {
	id, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("id = ", id)
	enti := storage.QueryOrder(id[0])
	order := &entity.Order{
		Uid: enti.Uid,
		Type: enti.OrderType,
		Content: enti.Content,
		HouseCountry: enti.HouseCountry,
		HouseProvince: enti.HouseProvince,
		HouseCity: enti.HouseCity,
		HouseAddress: enti.HouseAddress,
		HouseLayout: enti.HouseLayout,
		Price: enti.Price,
		Status: enti.Status,
		PlacerId: enti.PlacerId,
		AccepterId: enti.AccepterId,
		CreateTime: enti.CreateTime}
	rsp, err := json.Marshal(order)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func postOrderHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Order{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	id := storage.AddOrder(
		requestBody.Type,
		requestBody.Content,
		requestBody.HouseCountry,
		requestBody.HouseProvince,
		requestBody.HouseCity,
		requestBody.HouseAddress,
		requestBody.HouseLayout,
		requestBody.Price,
		requestBody.Status,
		requestBody.PlacerId,
		requestBody.AccepterId)

	ret := &entity.AddOrderResult{Id: id}

	rsp, err := json.Marshal(ret)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putOrderHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Order{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateOrder(
		requestBody.Uid,
		requestBody.Type,
		requestBody.Content,
		requestBody.HouseCountry,
		requestBody.HouseProvince,
		requestBody.HouseCity,
		requestBody.HouseAddress,
		requestBody.HouseLayout,
		requestBody.Price,
		requestBody.Status,
		requestBody.PlacerId,
		requestBody.AccepterId)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postOrderSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.OrderSearchByStatusPalcerId{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	if(0 == len(requestBody.PlacerId)){
		return
	}
	arr := storage.QueryOrdersByStatusPlacerId(requestBody.Length, requestBody.Offset, requestBody.Status, requestBody.PlacerId);

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))

}

func deleteOrderHandler(w http.ResponseWriter, r *http.Request)  {
	id, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no uid exist in url param")
	}
	fmt.Println("delete id : ", id);
	storage.DeleteOrderByUid(id[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

/*高级搜索*/
func postOrderSearchAdvancedHandler(w http.ResponseWriter, r *http.Request)  {

	qType, ok := r.URL.Query()["type"]
	if(!ok) {
		panic("no qType exist in url param")
	}
	fmt.Println("query type : ", qType);


	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	rspString := jsonUnmarshalQueryByType(qType[0], con)

	fmt.Println(rspString)

	CORSHandle(w)
	io.WriteString(w, rspString)

}

func jsonUnmarshalQueryByType(queryType string, bodyByte []byte) string{
	switch(queryType){
		case orderQueryTypeByIdCardNumber:
			return OrderQueryTypeByIdCardNumberHandler(bodyByte)
		case orderQueryTypeByPhoneNumber:
			return OrderQueryTypeByPhoneNumberHandler(bodyByte)
		case orderQueryTypeByRealName:
			return OrderQueryTypeByRealNameHandler(bodyByte)
		case orderQueryTypeBeforeTime:
			return OrderQueryTypeBeforeTimeHandler(bodyByte)
		case orderQueryTypeAfterTime:
			return OrderQueryTypeAfterTimeHandler(bodyByte)
		case orderQueryTypeRangeTime:
			return OrderQueryTypeRangeTimeHandler(bodyByte)
		case orderQueryTypeByStatusGroup:
			return OrderQueryTypeByStatusGroupHandler(bodyByte)
		case orderQueryTypeByAddress:
			return OrderQueryTypeByAddressHandler(bodyByte)
		case orderQueryTypeByLayoutGroup:
			return OrderQueryTypeByLayoutGroupHandler(bodyByte)
		case orderQueryTypeBelowPrice:
			return OrderQueryTypeBelowPriceHandler(bodyByte)
		case orderQueryTypeAbovePrice:
			return OrderQueryTypeAbovePriceHandler(bodyByte)
		case orderQueryTypeRangePrice:
			return OrderQueryTypeRangePriceHandler(bodyByte)
		case orderQueryTypeByOrderTypeGroup:
			return OrderQueryTypeByOrderTypeGroupHandler(bodyByte)
		case orderQueryTypeAll:
			return OrderQueryTypeAllHandler(bodyByte)
	}
	return ""
}

func OrderQueryTypeByIdCardNumberHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByIdCardNumber{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询用户ids*/
	users := storage.QueryUserByIdCardNumber(requestBody.PlacerIdCardNumber)
	if(0 == len(users)){
		fmt.Println("no user searched!")
		return ""
	}

	fmt.Println("fount users: ", users)

	userIds := make([]string, 0)
	for _,user := range users{
		userIds = append(userIds, user.Uid)
	}

	/*查询所有订单*/
	total, arr := storage.QueryOrderByUsers(userIds, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	qRet := &entity.OrderQueryResult{
		total,
		entities}

	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeByPhoneNumberHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByPhoneNumber{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询用户ids*/
	users := storage.QueryUserByPhoneNumber(requestBody.PlacerPhoneNumber)
	if(0 == len(users)){
		fmt.Println("no user searched!")
		return ""
	}

	fmt.Println("fount users: ", users)

	userIds := make([]string, 0)
	for _,user := range users{
		userIds = append(userIds, user.Uid)
	}

	/*查询所有订单*/
	total, arr := storage.QueryOrderByUsers(userIds, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	qRet := &entity.OrderQueryResult{
		total,
		entities}

	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeByRealNameHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByRealName{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询用户ids*/
	users := storage.QueryUserByName(requestBody.RealName)
	if(0 == len(users)){
		fmt.Println("no user searched!")
		return ""
	}

	fmt.Println("fount users: ", users)

	userIds := make([]string, 0)
	for _,user := range users{
		userIds = append(userIds, user.Uid)
	}

	/*查询所有订单*/
	total, arr := storage.QueryOrderByUsers(userIds, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeBeforeTimeHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryBeforeTime{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderBeforeTime(requestBody.Time, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeAfterTimeHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryAfterTime{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderAfterTime(requestBody.Time, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeRangeTimeHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryRangeTime{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderRangeTime(requestBody.FromTime, requestBody.ToTime, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeByStatusGroupHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByStatusGroup{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderByStatusGroup(requestBody.StatusGroup, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeByAddressHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByAddress{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total,arr := storage.QueryOrderByAddress(requestBody.Country,requestBody.Province,requestBody.City, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}
// QueryOrderByLayout

func OrderQueryTypeByLayoutGroupHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByLayoutGroup{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print("request body =", requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderByLayoutGroup(requestBody.LayoutGroup, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeBelowPriceHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryBelowPrice{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderBelowPrice(requestBody.Price, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeAbovePriceHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryAbovePrice{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderAbovePrice(requestBody.Price, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeRangePriceHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryRangePrice{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderRangePrice(requestBody.FromPrice, requestBody.ToPrice, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeByOrderTypeGroupHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryByOrderTypeGroup{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrderByOrderTypeGroup(requestBody.TypeGroup, requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}

	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}

func OrderQueryTypeAllHandler(bodyByte []byte) string{
	requestBody := &entity.OrderQueryAll{}

	err := json.Unmarshal(bodyByte, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	/*查询所有订单*/
	total, arr := storage.QueryOrders(requestBody.Offset, requestBody.Length)

	fmt.Println(arr)

	entities := make([]entity.Order, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Order{
			arr[i].Uid,
			arr[i].OrderType,
			arr[i].Content,
			arr[i].HouseCountry,
			arr[i].HouseProvince,
			arr[i].HouseCity,
			arr[i].HouseAddress,
			arr[i].HouseLayout,
			arr[i].Price,
			arr[i].Status,
			arr[i].PlacerId,
			arr[i].AccepterId,
			arr[i].CreateTime}
		entities = append(entities, *enti)
	}
	qRet := &entity.OrderQueryResult{
		total,
		entities}
	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))

	return string(rsp)
}