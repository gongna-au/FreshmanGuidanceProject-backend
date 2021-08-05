
package models

import (
	//"errors"
	"github.com/jinzhu/gorm"
	"fmt"
	"errors"
	DB "github.com/FreshmanGuidanceProject/api/database"

	//"github.com/jinzhu/gorm"
)
type Spot struct {
	Id           int  `json:"id"`           // 列名为 `id`
	Name         string `json:"name"`         // 列名为 `password`
	Introduction string `json:"introduction"` // 列名为 `num_of_spot`
}
func (spot Spot) GetSpotInformation() (s Spot, err error) {
	fmt.Println("GetSpotInformation()",spot)

	s,err=GetSpotInfoById(spot.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return s, errors.New("The requested landmark  doesn't exist !")
	}

	//根据id响应给前端信息
	return s, err

}
func GetSpotInfoById(id int) (s Spot, err error) {

	//初始化一个空的结构体
	s =Spot{}
	d := DB.Eloquent.Table("spots").Where("id = ?", id).First(&s)



	return s, d.Error
}


