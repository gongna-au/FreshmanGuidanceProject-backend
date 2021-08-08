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

// UpdateSpotAndKnowNum  操作用户移动更新用户探索数据接口
// @Summary UpdateSpotAndKnowNum 操作用户移动更新用户探索数据接口
// @Description 在URL中输入地标ID,Body里输入学号就可以实现更新用户个人的探索记录，不论是否探索过该地标都会得到地标的信息，实现了重复探索。
// @Accept application/json
// @Produce application/json
// @Param object body model.Person true "只输入学号"
// @Param id path int true "id"
// @Success 200 {object} model.Person model.Spot
// @Failure 400 {object} Response "{"Code":"10002", "Message":"Error occurred while binding the request body to the struct. "}
// @Failure 500 {object} Response "{"Code":"20002", "Message":"Database error:The landmark number you entered is incorrect, please try again ."}
// @Failure 500 {object} Response"{"Code":"20002", "Message":"Database error:The student ID you entered is incorrect, please try again ."}
// @Router /auth/spot/:id [put]
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
	//学号合法

	//判断学号前缀是否是2018/2019/2020/2021
	person, err = JudgeUserInputPreview(c, person)

	if (err != nil) || (person == model.Person{}) {

		SendError(c, err.Error())
		return

	}

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

// GetHonoraryTitle 获取荣誉称号接口
// @Summary GetHonoraryTitle 结束时获取荣誉称号接口接口
// @Description 输入学号就可以获得称号
// @Accept application/json
// @Produce application/json
// @Param object body model.Person true "输入学号就可以"
// @Success 200 {object} model.Person
// @Failure 400 {object} Response "{"Code":"10002", "Message":"Error occurred while binding the request body to the struct. "}
// @Failure 500 {object} Response "{"Code":"20002", "Message":"Database error ."}
// @Failure 500 {object} Response "{"Code":"20002", "Message":"Database error：The student ID you entered is incorrect, please try again ."}
// @Router /auth/spot/over [get]

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
	//判断解锁了多少个地标
	tag1, tag2, err := model.JudgeSpecialSpotNum(person.StudentId)

	JudgeUserLevel(c, person, tag1, tag2)
	//获取荣誉称号

}
