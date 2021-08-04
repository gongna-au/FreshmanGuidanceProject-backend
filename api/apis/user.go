package apis

import (
	"fmt"
	//DB "github.com/FreshmanGuidanceProject/api/database"
	//errno "github.com/FreshmanGuidanceProject/api/errno"
	model "github.com/FreshmanGuidanceProject-backend/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
	//"github.com/jinzhu/gorm"
)

type Person struct {
	ID        int    `json:"id"`          // 列名为 `id`
	StudentId string `json:"student_id"`  // 列名为 `student_id`
	Password  string `json:"password"`    // 列名为 `password`
	NumOfSpot int    `json:"num_of_spot"` // 列名为 `num_of_spot`
	NumOfKnow int    `json:"num_of_know"` // 列名为 `num_of_know"`

}

func Register(c *gin.Context) {

	var person model.Person

	err := c.BindJSON(&person)
	if err != nil {

		SendBadRequest(c, err, err.Error())

	}
	//判断输入是否合法
	person, err = JudgeUserInput(c, person)
	if (person == (model.Person{})) || (err != nil) {
		SendError(c, err.Error())
	}
	//调用功能函数
	u, err := person.Register()

	if err != nil {
		SendError(c, err.Error())
		return
	}
	SendResponse(c, u)
}

func Login(c *gin.Context) {

	var person model.Person

	err := c.BindJSON(&person)
	if err != nil {
		SendBadRequest(c, err, err.Error())

	}
	//调用功能函数
	u, err := person.Login()

	if err != nil {

		SendError(c, err.Error())
		return
	}
	SendResponse(c, u)

}

/*
func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := DB.Eloquent.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}*/
/*
func UpdatePerson(c *gin.Context) {

	var person Person
	id := c.Params.ByName("id")

	if err := DB.Eloquent.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	DB.Eloquent.Save(&person)
	c.JSON(200, person)

}
*/
/*
func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := DB.Eloquent.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}*/
/*
func GetPeople(c *gin.Context) {
	var people []Person
	if err := DB.Eloquent.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

}*/

func JudgeUserLevel(c *gin.Context, p model.Person) {

	u, err := model.GetUserByStudentId(p.StudentId)
	if (err != nil) || (u == model.Person{}) {
		SendError(c, err.Error())
		return
	}

	fmt.Println("JudgeUserLevel", u.NumOfKnow)
	var NumOfSpot float64
	var NumOfKnow float64
	NumOfSpot = float64(u.NumOfSpot)
	NumOfKnow = float64(u.NumOfKnow)

	var temp1 int
	var temp2 int
	var temp3 int
	temp1 = int((NumOfSpot / 14.0) * 100)
	temp2 = int((NumOfKnow / 21.0) * 100)
	temp3 = int(((NumOfSpot/14.0)*100 + (NumOfKnow/21.0)*100) / 2.0)

	OverPercentageOfSpot := "探索地标值" + strconv.Itoa(temp1) + "%"
	OverPercentageOfKnow := "获得见闻值" + strconv.Itoa(temp2) + "%"
	level := "超过了" + strconv.Itoa(temp3) + "%的山民"
	MountainNameNumber := "山民" + u.StudentId
	if (u.NumOfSpot <= 4) && (u.NumOfSpot >= 1) {

		SendResponse(c, u, MountainNameNumber, "Congratulations on getting the honorary title of ---Junior Explorer初级探索者---", OverPercentageOfSpot, OverPercentageOfKnow, level)
		return
	}
	if (u.NumOfSpot > 4) && (u.NumOfSpot <= 7) {

		SendResponse(c, u, MountainNameNumber, "Congratulations on winning the honorary title of ---Intermediate Explorer中级探索者---", OverPercentageOfSpot, OverPercentageOfKnow, level)
		return
	}
	if (u.NumOfSpot > 7) && (u.NumOfSpot <= 11) {

		SendResponse(c, u, MountainNameNumber, "Congratulations on winning the honorary title of ---Advanced Explorer进阶探索者---", OverPercentageOfSpot, OverPercentageOfKnow, level)
		return
	}

	if (u.NumOfSpot > 11) && (u.NumOfSpot < 14) {

		SendResponse(c, u, MountainNameNumber, "Congratulations on winning the honorary title of ---Pioneer Explorer先锋探索者---", OverPercentageOfSpot, OverPercentageOfKnow, level)

		return
	}

	if u.NumOfSpot == 14 {

		SendResponse(c, u, MountainNameNumber, "Congratulations on winning the honorary title of ---Pioneer Explorer先锋探索者---", OverPercentageOfSpot, OverPercentageOfKnow, level)

		return
	}

}
