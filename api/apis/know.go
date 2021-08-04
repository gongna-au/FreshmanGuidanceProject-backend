package apis

import (
	//"fmt"
	//DB "github.com/FreshmanGuidanceProject/api/database"
	//errno "github.com/FreshmanGuidanceProject/api/errno"
	model "github.com/FreshmanGuidanceProject-backend/api/models"
	"github.com/gin-gonic/gin"
	"strconv"
	//"github.com/jinzhu/gorm"
)

func UpdateSpotAndKnowNum(c *gin.Context) {
	//获取地标ID
	id, _ := strconv.Atoi(c.Param("id"))
	spot := model.Spot{Id: id}
	//先判断地标是否合法
	s, err := JudgeLandmarkIsLegal(c, spot)
	//地标合法s--最开始初始的地标，地标不合法s--空结构体
	if (err != nil) || (s == model.Spot{}) {

		SendError(c, err.Error())
		//地标不合法是返回，让前端重新请求
		return

	}

	//获取用户ID
	var person model.Person

	err = c.BindJSON(&person)

	if err != nil {

		SendBadRequest(c, err, err.Error())
		return

	}

	//判断输入的学号是否合法
	person, err = JudgeUserInput(c, person)

	if (err != nil) || (person == model.Person{}) {

		SendError(c, err.Error())
		return

	}
	//学号合法啊
	//判断用户是否存在
	u, err := model.GetUserByStudentId(person.StudentId)

	if (err != nil) || (u == model.Person{}) {
		//用户不存在
		SendError(c, err.Error())
		return

	}

	//调用功能函数
	s, person, err = spot.UpdateSpotAndKnowNum(person)

	if spot == (model.Spot{}) {
		SendError(c, err.Error())
		return
	}
	SendResponse(c, s, person)

}

func GetHonoraryTitle(c *gin.Context) {
	//获取用户ID
	var person model.Person

	err := c.BindJSON(&person)
	if err != nil {

		SendBadRequest(c, err, err.Error())

		return

	}
	//判断用户是否存在
	_, err = model.GetUserByStudentId(person.StudentId)
	if err != nil {

		SendError(c, err.Error())
		return

	}
	//判断输入是否合法
	person, err = JudgeUserInput(c, person)

	if (err != nil) || (person == model.Person{}) {

		SendError(c, err.Error())
		return

	}
	JudgeUserLevel(c, person)
	//获取荣誉称号

}
