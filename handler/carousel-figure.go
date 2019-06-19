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

func CarouselFigureHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getCarouselFigureHandler(w, r)
		break;
	case "POST":
		postCarouselFigureHandler(w, r)
		break;
	case "PUT":
		putCarouselFigureHandler(w, r)
		break;
	case "DELETE":
		deleteCarouselFigureHandler(w, r)
		break;
	case "OPTIONS":
		CORSHandle(w)
		break;
	}	
}

func CarouselFigureSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RequestURI);
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		break;
	case "POST":
		postCarouselFigureSearchHandler(w, r)
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

func postCarouselFigureHandler(w http.ResponseWriter, r *http.Request)  {

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.CarouselFigure{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	err = storage.AddCarouselFigure(
		requestBody.ImageUrl,
		requestBody.Location,
		requestBody.Desc)
	
	ret := ""
	if(nil != err){
		ret = GetErrJsonString(1, "错误: 添加出错!")
	}else{
		ret = GetSuccessJsonString()
	}

	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func getCarouselFigureHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	dbResult := storage.QueryCarouselFigure(uid[0])

	enti := entity.ConvertToCarouselFigureEntity(dbResult)

	rsp, err := json.Marshal(enti)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func putCarouselFigureHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.Body);

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.CarouselFigure{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	storage.UpdateCarouselFigure(
		requestBody.Uid,
		requestBody.ImageUrl,
		requestBody.Location,
		requestBody.Desc)

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func deleteCarouselFigureHandler(w http.ResponseWriter, r *http.Request)  {
	uid, ok := r.URL.Query()["uid"]
	if(!ok) {
		panic("no id exist in url param")
	}
	fmt.Println("uid = ", uid)
	storage.DeleteCarouselFigure(uid[0])

	ret := GetSuccessJsonString()
	fmt.Println(ret)
	CORSHandle(w)
	io.WriteString(w, ret)
}

func postCarouselFigureSearchHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("postCarouselFigureSearchHandler");

	con,_:=ioutil.ReadAll(r.Body)
	fmt.Println(string(con))

	requestBody := &entity.CarouselFigureRangeQuery{}
	err := json.Unmarshal(con, requestBody)
	Error.CheckErr(err)
	fmt.Print(requestBody)

	arr := storage.QueryCarouselFigureByRange(requestBody.Count, requestBody.Offset)

	entities := make([]entity.CarouselFigure, 0)
	for i := 0 ; i < len(arr) ; i ++ {
		enti := entity.ConvertToCarouselFigureEntity(&arr[i])
		entities = append(entities, *enti)
	}

	rsp, err := json.Marshal(entities)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}