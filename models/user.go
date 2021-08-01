package models

import (
	//"errors"
	//"github.com/jinzhu/gorm"
	"fmt"

	orm "github.com/FreshmanGuidanceProject/api/database"
)

type Person struct {
	ID        int64  `json:"id"`          // 列名为 `id`
	Username  string `json:"username"`    // 列名为 `username`
	Password  string `json:"password"`    // 列名为 `password`
	NumOfSpot int64  `json:"num_of_spot"` // 列名为 `num_of_spot`
	NumOfKnow int64  `json:"num_of_know"` // 列名为 `num_of_know"`

}

var Users []Person

func (user Person) CreatePerson() (id int64, err error) {

	//添加数据
	result := orm.Eloquent.Create(&user)

	id = user.ID
	if result.Error != nil {
		err = result.Error
		fmt.Println(err)
		return id, err
	} else {
		println("插入成功!")
	}
	return id, err
}

/*
func Register(username, password string) error {

	if username == "" || password == "" {
		return errors.New("lack for information")
	}
	//断用户是否存在
	user, err := GetUserByStuId(username)
	if errors.Is(err, gorm.ErrRecordNotFound) {

		//var errors []error
		user = &User{
			Username:       username,
			Password:       password,
			NumOfLandmarks: 0,
			NumOfKnowledge: 0,
		}
	}
	return user.CreateUserInDB()



}

func (u *User) CreateUserInDB() error {
	return orm.Eloquent.Create(u).Error
}

func GetUserByStuId(username string) (*User, error) {
	u := &User{}
	d := orm.Eloquent.Table("users").Where("username = ?", username).First(u)
	return u, d.Error
}
*/
/*
//检查错误是否为 RecordNotFound
//err := db.First(&user, 100).Error
//errors.Is(err, gorm.ErrRecordNotFound)
/*
//定义结构体切片
var Users []User

//添加
func (user User) Insert() (id int64, err error) {

	//添加数据
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//查找
func (user *User) Users() (users []User, err error) {

	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//修改
func (user *User) Update(id int64) (updateUser User, err error) {

	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
*/
