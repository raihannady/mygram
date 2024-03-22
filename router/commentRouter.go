package router

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.RouterGroup) {

    commentRouter := router.Group("/comments")
    {
        commentRouter.POST("/", controllers.CreateComment)
        commentRouter.GET("/", controllers.GetComment)
        commentRouter.GET("/:commentID", controllers.GetCommentByID)
        commentRouter.PUT("/:commentID", controllers.UpdateCommentRequest)
        commentRouter.DELETE("/:commentID", controllers.DeleteComment)
    }
}
