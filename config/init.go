package config

import(
	"fmt"
	"github.com/akkuman/parseConfig"
)

var(
	RunMode = "dev"
	DatabaseIp = "0.0.0.0:3306"
)


func Init() {
	var config = parseConfig.New("config.json")
	RunMode = config.Get("mode").(string)
	fmt.Println("config mode = ", RunMode);
	DatabaseIp = config.Get("database_ip").(string)
	fmt.Println("config database ip = ", DatabaseIp);
	ServerListentIp = config.Get("server_listen_ip").(string)
	fmt.Println("config server listen ip = ", ServerListentIp);
	
}
