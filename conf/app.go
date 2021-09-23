package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type appConfig struct {
	Server Server
}

type Server struct {
	PushAuthKey string
	WsAuthKey   string
	Addr        string
}

// 全局配置信息
var AppConf appConfig

func init() {
	appConfDir := "configs"
	appConfPath := appConfDir + "/app.toml"

	if !isFile(appConfPath) {
		log.Println("app配置：[" + appConfPath + "]不存在")

		config := strings.ReplaceAll(DefaultConfig, "{PushAuthKey}", randomStr(32))
		config = strings.ReplaceAll(config, "{WsAuthKey}", randomStr(32))

		createFile(appConfDir)
		tmpErr := ioutil.WriteFile(appConfPath, []byte(config), 0644)
		if tmpErr != nil {
			panic(tmpErr)
		}
		log.Println("app配置：[" + appConfPath + "]创建成功")
	}

	_, err := toml.DecodeFile(appConfPath, &AppConf)

	if err != nil {
		panic(err)
	}
	log.Println("App配置：" + fmt.Sprintf("%+v", AppConf))
}

func isFile(path string) bool {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return true
}

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func randomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(1000)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
