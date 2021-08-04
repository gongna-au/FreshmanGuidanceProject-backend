package main

import (
	//"fmt"

	DB "github.com/FreshmanGuidanceProject-backend/api/database"
	router "github.com/FreshmanGuidanceProject-backend/api/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //加载mysql
)

var err error

func main() {
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
