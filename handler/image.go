package handler

import(
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	// "mime"
	"mime/multipart"
	"os"
	"time"
	config "overseas-bulter-server/config"
	"encoding/json"
	// wx "overseas-bulter-server/wx"
	// storage "overseas-bulter-server/storage"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
)

// func HouseSearchHandler(w http.ResponseWriter, r *http.Request)  {
// 	fmt.Println("house search handler")
// 	fmt.Println(r.Method);
// 	switch(r.Method){
// 	case "GET":
// 		break;
// 	case "POST":
// 		postHouseSearchHandler(w, r)
// 		break;
// 	case "PUT":
// 		break;
// 	case "DELETE":
// 		break;
// 	case "OPTIONS":
// 		fmt.Println(r.Header.Get("Content-Type"))
// 		CORSHandle(w)
// 		// w.Header().Set("Access-Control-Allow-Origin", "*")
// 		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		// w.Header().Set("Content-Type", "application/json;charset=utf-8")
// 		break;
// 	}	
// }

func ImageHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("image handler")
	fmt.Println(r.Method);
	switch(r.Method){
	case "GET":
		getImageHandler(w, r)
		break;
	case "POST":
		postImageHandler(w, r)
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

func getImageHandler(w http.ResponseWriter, r *http.Request) {
	src, ok := r.URL.Query()["src"]
	if(!ok) {
		panic("no src exist in url param")
	}
	fmt.Println("src = ", src)
    http.ServeFile(w, r, src[0])
}

func postImageHandler(w http.ResponseWriter, r *http.Request)  {

	fileContent := parseMultiPart(r)

	filePath := writeFile(fileContent);

	rspString := config.ImageServerAddress + config.GenerateIntegratedUri("/image") + "?src=" + filePath

	fmt.Println(rspString)

	entity := &entity.ImageUploadResult{Path: rspString}
	rsp, err := json.Marshal(entity)
	Error.CheckErr(err)
	fmt.Print(string(rsp))
	CORSHandle(w)
	io.WriteString(w, string(rsp))
}

func parseMultiPart(r *http.Request)  []byte{
	mr,err := r.MultipartReader()
	Error.CheckErr(err)
	form ,_ := mr.ReadForm(128)
	return getFormData(form)
}

func getFormData(form *multipart.Form) []byte{

	//获取 multi-part/form body中的form value
	
	for k,v := range form.Value{
		fmt.Println("value,k,v = ",k,",",v)
	}
	
	fmt.Println()
	
	//获取 multi-part/form中的文件数据
	
	for _,v := range form.File {
		for i:=0 ;i < len(v);i++{
		
		fmt.Println("file part ",i,"-->")
		fmt.Println("fileName   :",v[i].Filename)
		fmt.Println("part-header:",v[i].Header)
		fmt.Println("part-header-0:",v[i].Header["Content-Disposition"])
		fmt.Println("part-header-1:",v[i].Header["Content-Type"])

		f,_ := v[i].Open()
		buf,_:= ioutil.ReadAll(f)
		fmt.Println("file-content-length",len(string(buf)))
		return buf

		}
	}
	return nil
	
}

func writeFile(content []byte) string{
	
	currentTime := time.Now()
	timeStamp := currentTime.Unix()

	todayString := fmt.Sprintf("%04d%02d%02d", currentTime.Year(), currentTime.Month(), currentTime.Day())

	todayImageDir := "image-upload/" + todayString
	imageName := fmt.Sprintf("%d.jpg", timeStamp)
	imageFullName := todayImageDir + "/"+ imageName

	fmt.Println(imageFullName);

	err := os.MkdirAll(todayImageDir, 0777)
	Error.CheckErr(err)
	err2 := ioutil.WriteFile(imageFullName , content, 0777) //写入文件(字节数组)
	Error.CheckErr(err2)
	
	return imageFullName
}