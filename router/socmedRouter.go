package router

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func SocmedRouter(router *gin.RouterGroup) {
	socmedRouter := router.Group("/socialmedias")
	{
		socmedRouter.POST("/", controllers.CreateSocialMediaRequest)
		socmedRouter.GET("/", controllers.GetSocialMedia)
		socmedRouter.GET("/:socialMediaID", controllers.GetSocialMediaByID)
		socmedRouter.PUT("/:socialMediaID", controllers.UpdateSocialMediaRequest)
		socmedRouter.DELETE("/:socialMediaID", controllers.DeleteSocialMedia)
	}
}