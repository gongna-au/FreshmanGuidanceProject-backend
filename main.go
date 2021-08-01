package main

import (
	"fmt"

	model "github.com/FreshmanGuidanceProject/api/models"

	//apis"github.com/FreshmanGuidanceProject/api/apis"
	"github.com/gin-gonic/gin"
	//"github.com/FreshmanGuidanceProject/api/router"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type Person struct {
	ID        int64  `json:"id"`          // 列名为 `id`
	StudentId string `json:"student_id"`  // 列名为 `student_id`
	Password  string `json:"password"`    // 列名为 `password`
	NumOfSpot int64  `json:"num_of_spot"` // 列名为 `num_of_spot`
	NumOfKnow int64  `json:"num_of_know"` // 列名为 `num_of_know"`

}

func main() {

	db, err = gorm.Open("mysql", "gongna2:123456@tcp/Test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/people/", GetPeople)
	r.GET("/people/:id", GetPerson)
	r.POST("/people", CreatePerson)
	r.PUT("/people/:id", UpdatePerson)
	r.DELETE("/people/:id", DeletePerson)
	//r=router.InitRouter(r)

	r.Run(":8080")
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePerson(c *gin.Context) {

	var person Person
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)

}

/*
func CreatePerson(c *gin.Context) {

	var person Person
	c.BindJSON(&person)

	db.Create(&person)
	c.JSON(200, person)
}

*/
func CreatePerson(c *gin.Context) {

	var person model.Person
	err := c.BindJSON(&person)
	if err != nil {
		println("c.BindJSON失败")
		println(err)
	} else {
		println("c.BindJSON成功")
	}

	//db.Create(&person)
	_, err = person.CreatePerson()
	if err != nil {
		println("调用功能函数失败")
		println(err)
		c.JSON(400, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	} else {
		println("调用功能函数成功")
	}
	c.JSON(200, person)
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}
func GetPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}

/*
package main

import (
	"fmt"
	_ "github.com/FreshmanGuidanceProject/api/database"
	model "github.com/FreshmanGuidanceProject/api/models"
	//orm "github.com/FreshmanGuidanceProject/api/database"
	"github.com/FreshmanGuidanceProject/api/router"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB
var err error

func main() {

	Eloquent, err = gorm.Open("mysql", "gongna2:123456@tcp/FreshmanGuidance?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer Eloquent.Close()
	Eloquent.AutoMigrate(&model.User{})
	//defer orm.Eloquent.Close()
	//router := router.InitRouter()
	router := router.InitRouter()
	router.Run(":8000")
}
*/
