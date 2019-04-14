package config

import(
	"fmt"
)

const ServerListentIp = "127.0.0.1:8008"//"0.0.0.0:8008"//
const preHttpString = "/overseas-bulter/v1"


func GenerateIntegratedUri(shortcut_api string) string {
	result := preHttpString + shortcut_api
	fmt.Println(result)
	return result
}