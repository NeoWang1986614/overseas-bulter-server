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

func EmployeeCheckHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("employee handler")
	fmt.Println(r.Method);

	if("POST" == r.Method){
		fmt.Println(r.Body);
		con,_:=ioutil.ReadAll(r.Body)
		fmt.Println(string(con))

		requestBody := &entity.EmployeeCheck{}
		err := json.Unmarshal(con, requestBody)
		Error.CheckErr(err)
		fmt.Print(requestBody)

		employeeArr := storage.QueryEmployeeByUserName(requestBody.UserName)
		result := ""
		fmt.Println("employeeArr", employeeArr)
		if(0 == len(employeeArr)){
			result = GetErrJsonString(1001, "user name not found in db")
		}else{
			if(employeeArr[0].Password != requestBody.Password){
				result = GetErrJsonString(1002, "password incorrect")
			}else{
				result = GetErrJsonString(0, "check success")
			}
		}

		CORSHandle(w)
		io.WriteString(w, result)
	} else if ("OPTIONS" == r.Method){
		CORSHandle(w)
	}
}

func EmployeeHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("employee handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getEmployeeHandler(w, r)
		break;
	case "POST":
		postEmployeeHandler(w, r)
		break;
	case "PUT":
		putEmployeeHandler(w, r)
		break;
	case "DELETE":
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func EmployeeSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("empoyee search handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postEmployeeSearchHandler(w, r)
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


func getEmployeeHandler(w http.ResponseWriter, r *http.Request)  {
	eid, ok := r.URL.Query()["id"]
	if(!ok) {
		panic("no eid exist in url param")
	}

	enti := storage.QueryEmployee(eid[0])
	employee := &entity.Employee{
		Id: enti.Id,
		UserName: enti.UserName,
		Password: enti.Password,
		Nickname: enti.Nickname,
		AvatarUrl: enti.AvatarUrl,
		PhoneNumber: enti.PhoneNumber}

	rsp, err := json.Marshal(employee)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putEmployeeHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Employee{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateEmployee(
		requestBody.Id,
		requestBody.UserName,
		requestBody.Password,
		requestBody.Nickname,
		requestBody.AvatarUrl,
		requestBody.PhoneNumber)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postEmployeeHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.Employee{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.AddEmployee(
		requestBody.UserName,
		requestBody.Password,
		requestBody.Nickname,
		requestBody.AvatarUrl,
		requestBody.PhoneNumber)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postEmployeeSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.QueryEmployeeRange{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)
	
	total, arr := storage.QueryEmployees(requestBody.Offset, requestBody.Length);
	
	
	entities := make([]entity.Employee, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		var enti = &entity.Employee{
			arr[i].Id,
			arr[i].UserName,
			arr[i].Password,
			arr[i].Nickname,
			arr[i].AvatarUrl,
			arr[i].PhoneNumber}
		entities = append(entities, *enti)
	}

	qRet := &entity.QueryEmployeeResult{
		total,
		entities}

	rsp, err := json.Marshal(qRet)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}
