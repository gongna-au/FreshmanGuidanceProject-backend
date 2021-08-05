package router

import (
	apis "github.com/FreshmanGuidanceProject/api/apis"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadRouter(r *gin.Engine) *gin.Engine {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authRouter := r.Group("/api/v1/auth")
	{
		authRouter.POST("/register", apis.Register)
		authRouter.POST("/login", apis.Login)
		//authRouter.GET("/people/", apis.GetPeople)
		//authRouter.GET("/people/:id", apis.GetPerson)
		authRouter.PUT("/spot/:id", apis.UpdateSpotAndKnowNum)

	}
	spotRouter := r.Group("/api/v1/spot")
	{
		//spotRouter.GET("/:id", apis.GetSpotInformation)
		spotRouter.GET("/over", apis.GetHonoraryTitle)

	}
	return r

}
