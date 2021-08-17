package database

import (
	"fmt"
	"github.com/FreshmanGuidanceProject/setting"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"log"
)

var Eloquent *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	//Eloquent, err = gorm.Open("mysql", "root:123456@tcp/Test?charset=utf8")
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	//if Eloquent.Error != nil {
	//	fmt.Printf("database error %v", Eloquent.Error)
	//}
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	Eloquent, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	Eloquent.DB().SetMaxIdleConns(10000)
	Eloquent.DB().SetMaxOpenConns(10000)
}

func CloseDB() {
	defer Eloquent.Close()
}
