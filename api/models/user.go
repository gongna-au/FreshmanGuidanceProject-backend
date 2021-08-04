package models

import (
	//"errors"
	//"github.com/jinzhu/gorm"
	"errors"
	//"fmt"
	DB "github.com/FreshmanGuidanceProject-backend/api/database"

	"github.com/jinzhu/gorm"
)

type Person struct {
	ID        int    `json:"id"`          // 列名为 `id`
	StudentId string `json:"student_id"`  // 列名为 `student_id`
	Password  string `json:"password"`    // 列名为 `password`
	NumOfSpot int    `json:"num_of_spot"` // 列名为 `num_of_spot`
	NumOfKnow int    `json:"num_of_know"` // 列名为 `num_of_know"`

}

func (user Person) Register() (u Person, err error) {
	//初始化一个空结构体，用户不存在时返回
	a := Person{}
	if user.Password == "" {
		return a, errors.New("Missing user password information")
	}

	//判断用户是否存在
	u, err = GetUserByStudentId(user.StudentId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//用户没有找到
		//往数据库插入数据
		d := DB.Eloquent.Table("persons").Create(&user)

		return user, d.Error

	} else {
		//用户找到
		//返回信息错误
		return a, errors.New("user had existed")

	}

}

func (user Person) Login() (u Person, err error) {
	//先查询用户是否存在
	u, err = GetUserByStudentIdAndPassword(user.StudentId, user.Password)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//用户不存在
		return u, errors.New("user doesn't exist !")
	}

	return u, err

}

func GetUserByStudentIdAndPassword(studentId string, password string) (u Person, err error) {
	u = Person{}

	d := DB.Eloquent.Table("persons").Where("student_id = ? AND password = ?", studentId, password).First(&u)

	return u, d.Error
}

func GetUserByStudentId(studentId string) (u Person, err error) {
	u = Person{}

	d := DB.Eloquent.Table("persons").Where("student_id = ?", studentId).First(&u)

	return u, d.Error
}
