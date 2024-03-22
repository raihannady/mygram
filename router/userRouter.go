package router

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userID", controllers.UpdateUser)
		userRouter.DELETE("/:userID", controllers.DeleteUser)
	}
}