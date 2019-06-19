package config

import(
	"fmt"
)

var ServerListentIp = "0.0.0.0:8008"
const ImageServerAddress = "https://bulter.mroom.com.cn:8008"
const FileServerAddress = "https://bulter.mroom.com.cn:8008"
const preHttpString = "/overseas-bulter/v1"


func GenerateIntegratedUri(shortcut_api string) string {
	result := preHttpString + shortcut_api
	fmt.Println(result)
	return result
}
