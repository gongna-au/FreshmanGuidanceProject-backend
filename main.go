package main

import (
	//"fmt"

	DB "github.com/FreshmanGuidanceProject/api/database"
	router "github.com/FreshmanGuidanceProject/api/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //加载mysql
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

	//数据库的初始化
	DB.Init()
	defer DB.Eloquent.Close()
	//database.Eloquent, err = gorm.Open("mysql", "gongna2:123456@tcp/Test?charset=utf8&parseTime=True&loc=Local")

	//新建路由
	r := gin.Default()
	//加载路由
	r = router.LoadRouter(r)
	r.Run(":8080")
}
