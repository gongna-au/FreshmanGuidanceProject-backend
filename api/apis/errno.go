package apis

import (
	//"errors"
	"errors"
	//"fmt"
	errno "github.com/FreshmanGuidanceProject-backend/api/errno"
	model "github.com/FreshmanGuidanceProject-backend/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	//"github.com/jinzhu/gorm"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, u ...interface{}) {

	c.JSON(http.StatusOK, u)

}
func SendBadRequest(c *gin.Context, err error, cause string) {

	c.JSON(http.StatusBadRequest, Response{
		Code:    errno.ErrBind.Code,
		Message: errno.ErrBind.Message + ": " + cause,
	})
}

func SendError(c *gin.Context, cause string) {

	c.JSON(http.StatusInternalServerError, Response{
		Code:    errno.ErrDatabase.Code,
		Message: errno.ErrDatabase.Message + ": " + cause,
	})
}

func SendlandmarkError(c *gin.Context, err error, cause string) {

	c.JSON(http.StatusInternalServerError, Response{
		Code:    errno.ErrDatabase.Code,
		Message: errno.ErrDatabase.Message + ": " + cause,
	})

}
func SendVerificationFiled(c *gin.Context, err error, cause string) {

	c.JSON(http.StatusBadRequest, Response{
		Code:    errno.ErrValidation.Code,
		Message: errno.ErrValidation.Message + ": " + cause,
	})

}

func JudgeUserInput(c *gin.Context, person model.Person) (u model.Person, err error) {

	u = model.Person{}
	//判断学号是否有非法字符
	_, err = strconv.Atoi(person.StudentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errno.ErrUserInput.Code,
			Message: errno.ErrUserInput.Message + "More specific information is：" + "The student ID entered by the user contains illegal characters",
			Data:    person,
		})
		//有非法字符，返回空结构体
		return u, errors.New("The student ID you entered is incorrect, please try again")

	}
	var i int = 0
	for range person.StudentId {
		i++
	}
	//判断学号是否是10位
	if i != 10 {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errno.ErrUserInput.Code,
			Message: errno.ErrUserInput.Message + "More specific information is：" + "The length of the student ID entered by the user is incorrect",
			Data:    person,
		})
		//不是10位学号，返回空结构体
		return u, errors.New("The student ID you entered is incorrect, please try again")

	}
	//如果是合法的学号，返回传入的结构体和nil
	return person, nil

}

func JudgeLandmarkIsLegal(c *gin.Context, spot model.Spot) (s model.Spot, err error) {
	s = model.Spot{}
	if spot.Id > 14 || spot.Id < 1 {
		c.JSON(http.StatusBadRequest, Response{
			Code:    errno.ErrUserInput.Code,
			Message: errno.ErrUserInput.Message + "More specific information is：" + "The number of landmarks entered is invalid",
		})
		//地标不合法返回空spot结构体
		return s, errors.New("The landmark number you entered is incorrect, please try again")
	}
	if spot.Id >= 1 && spot.Id <= 7 {
		c.JSON(http.StatusOK, Response{
			Message: "This landmark is a special landmark. Congratulations on unlocking a special landmark~",
		})
		//地标合法返回初始的地标
		return spot, nil
	}
	if spot.Id > 7 && spot.Id <= 14 {
		c.JSON(http.StatusOK, Response{

			Message: "This landmark is a common landmark. Congratulations on unlocking one of the common landmarks~",
		})
		//地标合法返回初始的地标
		return spot, nil
	}

	return spot, nil
}
