
package handler

import(
	"fmt"
	"net/http"
	"encoding/json"
	entity "overseas-bulter-server/entity"
	Error "overseas-bulter-server/error"
)

func GetErrJsonString(code uint, message string) string{
	ret, err := json.Marshal(entity.GetErr(code, message))
	Error.CheckErr(err)
	fmt.Print(string(ret))
	return string(ret)
}

func GetSuccessJsonString() string{
	ret, err := json.Marshal(entity.GetErrForSuccess())
	Error.CheckErr(err)
	fmt.Print(string(ret))
	return string(ret)
}

func CORSHandle(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
}