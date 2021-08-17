package config

import (
	//"fmt"

	//"fmt"
	"fmt"
	"log"

	"github.com/spf13/viper"
	//"log"
	"time"
)

var (
	V            *viper.Viper
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
)

func Init() {

	V = viper.New()

	V.SetConfigName("config.yaml")

	V.AddConfigPath("./config/")

	V.SetConfigType("yaml")

	if err := V.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Print("配置文件未找到")
			// 配置文件未找到错误；如果需要可以忽略
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Fatal(err)
		}

	}
	LoadBase()
	LoadServer()
	LoadApp()

}
func LoadBase() {
	RunMode = V.GetString("RUN_MODE")

}
func LoadServer() {
	HTTPPort = V.GetInt("server.HTTP_PORT")
	ReadTimeout = time.Duration((V.GetInt64("server.ReadTimeout"))) * time.Second
	WriteTimeout = time.Duration(V.GetInt64("server.WriteTimeout")) * time.Second
}
func LoadApp() {

	PageSize = V.GetInt("PAGE_SIZE")
}
