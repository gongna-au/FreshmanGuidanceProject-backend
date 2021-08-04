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

type Spot struct {
	Id           int    `json:"id"`           // 列名为 `id`
	Name         string `json:"name"`         // 列名为 `password`
	Introduction string `json:"introduction"` // 列名为 `num_of_spot`
}

func GetSpotInformation(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	spot := model.Spot{Id: id}

	//调用功能函数
	spot, err := spot.GetSpotInformation()
	if err != nil {
		SendError(c, err.Error())
		return
	}
	SendResponse(c, spot)

}
