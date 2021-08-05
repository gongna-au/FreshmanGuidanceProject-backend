package models

import (
	//"errors"
	//apis "github.com/FreshmanGuidanceProject/api/apis"
	"errors"
	"fmt"
	DB "github.com/FreshmanGuidanceProject/api/database"
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm"
)

type Exploration struct {
	ID        int    `json:"id"`         // 列名为 `id`
	StudentId string `json:"student_id"` // 列名为 `student_id`
	SpotId    int    `json:"spot_id"`    // 列名为 `spot_id"`
}

func (spot Spot) UpdateSpotAndKnowNum(u Person) (s Spot, p Person, d error) {
	//这里的用户一定是存在且合法
	//这里的地标一定是合法的
	s = Spot{}
	p = Person{}

	e := Exploration{
		StudentId: u.StudentId,
		SpotId:    spot.Id,
	}

	//先判断地标类型tag1==2特殊地标  tag1==1普通地标
	tag1 := JudgelandMarkRange(spot.Id)

	//根据地标响应信息

	s, _ = GetSpotInfoById(spot.Id)
	//s是我们无论是否探索过该地标都要响应的信息

	//判断是否探索过该地标tag2
	tag2 := SearchIsExplored(u.StudentId, spot.Id)

	if tag2 {
		//已经探索过该地标，还是可以返回用户信息
		u, _ = GetUserByStudentId(u.StudentId)
		return s, u, nil
	}
	if (tag1 == 1) && (!tag2) {

		//没有探索过并且是普通地标
		//新增见闻+1
		u, _ = UpdateSpotById(u.StudentId)
		u, _ = UpdateKnowById(u, tag1)
		//新增探索记录
		e, _ = e.CreateExplorationRecord()
		return s, u, nil

	}
	if (tag1 == 2) && (!tag2) {

		//新增见闻+2
		u, _ = UpdateSpotById(u.StudentId)
		u, _ = UpdateKnowById(u, tag1)

		//新增探索记录
		e, _ = e.CreateExplorationRecord()
		return s, u, nil

	}

	//根据id响应给前端信息
	return

}
func JudgeLandmarkIsLegal(spot Spot, u Person) (tag int) {
	tag1 := SearchIsExplored(u.StudentId, spot.Id)
	if spot.Id > 14 || spot.Id < 1 || tag1 == false {

		return 0
	}
	if spot.Id >= 1 && spot.Id <= 7 && tag1 == true {

		return 2
	}
	if spot.Id > 7 && spot.Id <= 14 && tag1 == true {

		return 1
	}

	return 0
}

func SearchIsExplored(studentId string, spotId int) bool {

	e := Exploration{}
	err := DB.Eloquent.Table("explorations").Where("student_id = ? And spot_id = ?", studentId, spotId).First(&e).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(errors.Is(err, gorm.ErrRecordNotFound))
		//探索过返回true
		return false
	}

	//没有探索过返回false
	return true

}
func JudgelandMarkRange(spotId int) (tag int) {

	if (spotId >= 1) && (spotId <= 7) {
		return 2
	}
	if (spotId >= 8) && (spotId <= 14) {
		return 1
	}

	return

}

func (exploration Exploration) CreateExplorationRecord() (e Exploration, err error) {

	s := DB.Eloquent.Table("explorations").Create(&exploration)
	//根据id响应给前端信息
	return exploration, s.Error

}

func UpdateSpotById(studentId string) (u Person, err error) {

	//地标数+1
	d := DB.Eloquent.Table("persons").Where("student_id = ?", studentId).Update("num_of_spot", gorm.Expr("num_of_spot * ?+ ?", 1, 1))
	//见闻数+2
	//d = DB.Eloquent.Table("persons").Where("student_id = ?", studentId).Update("num_of_know", gorm.Expr("num_of_spot * ?+ ?", 1, 1))

	//查寻用户更新后的信息
	DB.Eloquent.Table("persons").Where("student_id = ?", studentId).Find(&u)

	return u, d.Error
}
func UpdateKnowById(p Person, tag int) (u Person, err error) {

	//地标数+1
	//d := DB.Eloquent.Table("persons").Where("student_id = ?", studentId).Update("num_of_spot", gorm.Expr("num_of_spot * ?+ ?", 1, 1))
	//见闻数+2
	d := DB.Eloquent.Table("persons").Where("student_id = ?", p.StudentId).UpdateColumn("num_of_know", p.NumOfKnow+tag)

	//查寻用户更新后的信息
	DB.Eloquent.Table("persons").Where("student_id = ?", p.StudentId).Find(&u)

	return u, d.Error
}
func UpdateKnowById2(p Person) (u Person, err error) {

	//地标数+1
	//d := DB.Eloquent.Table("persons").Where("student_id = ?", studentId).Update("num_of_spot", gorm.Expr("num_of_spot * ?+ ?", 1, 1))
	//见闻数+2
	d := DB.Eloquent.Table("persons").Where("student_id = ?", p.StudentId).UpdateColumn("num_of_know", p.NumOfKnow+2)

	//查寻用户更新后的信息
	DB.Eloquent.Table("persons").Where("student_id = ?", p.StudentId).Find(&u)

	return u, d.Error
}

func GetSpotAndKnowNum(studentId string) (u Person, err error) {

	//查询最新的用户的探索记录
	d := DB.Eloquent.Table("persons").Where("student_id = ?", studentId).First(&u)

	return u, d.Error
}
func GetUserInfoById(studentId string) (u Person, err error) {

	//查询最新的用户的探索记录
	d := DB.Eloquent.Table("persons").Where("student_id = ?", studentId).First(&u)

	return u, d.Error
}

//判断解锁了多少个地标
func JudgeSpecialSpotNum(studentId string) (specialSpot1 int, specialSpot2 int, err error) {
	rows, err := DB.Eloquent.Table("explorations").Where("student_id = ?", studentId).Rows()

	defer rows.Close()
	SpecialLandmarks := []int{1, 2, 3, 4, 5, 6, 7}
	CommenLandmarks := []int{8, 9, 10, 11, 12, 13, 14}
	//tag1是特殊地标数
	tag1 := 0
	for _, v := range SpecialLandmarks {

		for rows.Next() {
			var exploration Exploration
			// ScanRows 方法用于将一行记录扫描至结构体
			DB.Eloquent.ScanRows(rows, &exploration)
			// 获取exploration的地标ID
			if exploration.SpotId == v {

				tag1++
				break
			}

		}
	}
	tag2 := 0
	//tag2是普通地标数
	for _, v := range CommenLandmarks {

		for rows.Next() {
			var exploration Exploration
			// ScanRows 方法用于将一行记录扫描至结构体
			DB.Eloquent.ScanRows(rows, &exploration)
			// 获取exploration的地标ID
			if exploration.SpotId == v {

				tag2++
				break
			}

		}
	}

	return tag1, tag2, err

}
