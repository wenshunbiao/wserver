package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type appConfig struct {
	Server Server
}

type Server struct {
	PushAuthToken string
	Addr          string
}

// 全局配置信息
var AppConf appConfig

func init() {
	_, err := toml.DecodeFile("configs/app.toml", &AppConf)
	if err != nil {
		panic(err)
	}
	log.Println("App配置：" + fmt.Sprintf("%+v", AppConf))
}
