package apis

import (
	//"fmt"
	//DB "github.com/FreshmanGuidanceProject/api/database"
	//errno "github.com/FreshmanGuidanceProject/api/errno"
	model "github.com/FreshmanGuidanceProject/api/models"
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

// Register 用户注册接口
// @Summary Register 用户注册接口
// @Description 可通过10位学号和密码注册账号(前缀必须是2018、2019、2020，2021)
// @Accept application/json
// @Produce application/json
// @Param object body model.Person true "登录的用户信息"
// @Success 200 {object} model.Person
// @Failure 400 {object} Response "{"Code":"10002", "Message":"Error occurred while binding the request body to the struct."}
// @Failure 500 {object} Response "{"Code":"20002", "Message":"Database error:Missing user password information"}
// @Failure 500 {object} Response "{"Code":"20002", "Message":"Database error:User had existed"}
// @Router /auth/register [post]

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

	//判断学号前缀是否是2018/2019/2020/2021
	person, err = JudgeUserInputPreview(c, person)

	if (err != nil) || (person == model.Person{}) {

		SendError(c, err.Error())
		return

	}
	//调用功能函数
	u, err := person.Register()

	if err != nil {
		SendError(c, err.Error())
		return
	}
	SendResponse(c, u)
}

// Register 用户登录接口
// @Summary Login 用户登录接口
// @Description 输入学号和密码登录
// @Accept application/json
// @Produce application/json
// @Param object body model.Person true "登录的用户信息"
// @Success 200 {object} model.Person
// @Failure 400 {object} Response "{"Code":"10002", "Message":"Error occurred while binding the request body to the struct. "}
// @Failure 500 {object} Response "{"Code":"20002", "Message":"Database error:The student ID you entered is incorrect, please try again"}
// @Router /auth/login [post]
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

func JudgeUserLevel(c *gin.Context, p model.Person, tag1 int, tag2 int) {

	u, err := model.GetUserByStudentId(p.StudentId)
	if (err != nil) || (u == model.Person{}) {
		SendError(c, err.Error())
		return
	}

	var NumOfSpot float64
	var NumOfKnow float64
	NumOfSpot = float64(u.NumOfSpot)
	NumOfKnow = float64(u.NumOfKnow)

	var temp1 int
	var temp2 int
	//var temp3 int
	temp1 = int((NumOfSpot / 14.0) * 100)
	temp2 = int((NumOfKnow / 21.0) * 100)
	//temp3 = int(((NumOfSpot/14.0)*100 + (NumOfKnow/21.0)*100) / 2.0)

	OverPercentageOfSpot := "探索地标值" + strconv.Itoa(temp1) + "%"
	OverPercentageOfKnow := "获得见闻值" + strconv.Itoa(temp2) + "%"
	//level := "超过了" + strconv.Itoa(temp3) + "%的山民"
	MountainNameNumber := "山民" + u.StudentId

	//探索的特殊地标数
	NumOfSpeciaLandmarks := "点亮特殊地标：" + strconv.Itoa(tag1)
	//探索的普通地标数
	NumOfCommenLandmarks := "点亮普通地标：" + strconv.Itoa(tag2)
	//判断解锁了多少个特殊地标

	theFirstexplorer := "成为第" + strconv.Itoa(u.ID) + "个点亮所有地标的探索者"
	if (u.NumOfKnow < 8) && (u.NumOfKnow >= 0) {

		SendResponse(c, u, MountainNameNumber, "你在本次登岛探索中", NumOfSpeciaLandmarks, NumOfCommenLandmarks, OverPercentageOfKnow, OverPercentageOfSpot, "获得称号--初级探索者")
		return
	}
	if (u.NumOfKnow > 7) && (u.NumOfKnow < 15) {

		SendResponse(c, u, MountainNameNumber, "你在本次登岛探索中", NumOfSpeciaLandmarks, NumOfCommenLandmarks, OverPercentageOfKnow, OverPercentageOfSpot, "获得称号--进阶探索者")
		return
	}

	if (u.NumOfKnow > 14) && (u.NumOfKnow < 21) {

		SendResponse(c, u, MountainNameNumber, "你在本次登岛探索中", NumOfSpeciaLandmarks, NumOfCommenLandmarks, OverPercentageOfKnow, OverPercentageOfSpot, "获得称号--先锋探索者")
		return
	}

	if u.NumOfKnow == 21 {

		SendResponse(c, u, MountainNameNumber, "你在本次登岛探索中", NumOfSpeciaLandmarks, NumOfCommenLandmarks, OverPercentageOfKnow, theFirstexplorer, "获得称号--先锋探索者")

		return
	}

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
