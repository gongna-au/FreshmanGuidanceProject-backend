package main

import (
	"fmt"
	//"log"
	"net/http"

	DB "github.com/FreshmanGuidanceProject/api/database"
	router "github.com/FreshmanGuidanceProject/api/router"
	config "github.com/FreshmanGuidanceProject/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// @title 新生引导项目
	// @version 1.0
	// @description 新生引导项目API
	// @termsOfService http://swagger.io/terms/
	// @contact.name gongna
	// @contact.url http://www.swagger.io/support
	// @contact.email 20364719155@qq.com
	// @license.name Apache 2.0
	// @license.url
	// @host 8080
	// @BasePath /api/v1

	//配置文件的初始化
	config.Init()

	dbType := config.V.GetString("db.type")

	user := config.V.GetString("db.user")
	fmt.Println(config.V.GetString("db.user"))
	password := config.V.GetString("db.password")
	fmt.Println(config.V.GetString("db.password"))
	host := config.V.GetString("db.host")
	dbName := config.V.GetString("db.database")
	DB.Eloquent, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		fmt.Println(err)
	}
	DB.Eloquent.DB().SetMaxIdleConns(10000)
	DB.Eloquent.DB().SetMaxOpenConns(10000)

	//数据库的初始化
	//DB.Init()

	//新建路由
	r := gin.Default()

	//加载路由
	r = router.LoadRouter(r)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.HTTPPort),
		Handler:      r,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}
	s.ListenAndServe()
}
