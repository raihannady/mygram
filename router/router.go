package router

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    addCommentRoutes(router)
    addPhotosRoutes(router)
    addSocmedRoutes(router)
    addUserRoutes(router)

    return router
}

func addCommentRoutes(r *gin.Engine) {
    commentRouter := r.Group("/comments")
    {
        commentRouter.POST("/", controllers.CreateComment)
        commentRouter.GET("/", controllers.GetComment)
        commentRouter.GET("/:commentID", controllers.GetCommentByID)
        commentRouter.PUT("/:commentID", controllers.UpdateCommentRequest)
        commentRouter.DELETE("/:commentID", controllers.DeleteComment)
    }
}

func addPhotosRoutes(r *gin.Engine) {
    photosRouter := r.Group("/photos")
    {
        photosRouter.POST("/", controllers.CreatePhoto)
        photosRouter.GET("/", controllers.GetPhotos)
        photosRouter.GET("/:photoID", controllers.GetPhotosByID)
        photosRouter.PUT("/:photoID", controllers.UpdatePhoto)
        photosRouter.DELETE("/:photoID", controllers.DeletePhoto)
    }
}

func addSocmedRoutes(r *gin.Engine) {
    socmedRouter := r.Group("/socialmedias")
    {
        socmedRouter.POST("/", controllers.CreateSocialMediaRequest)
        socmedRouter.GET("/", controllers.GetSocialMedia)
        socmedRouter.GET("/:socialMediaID", controllers.GetSocialMediaByID)
        socmedRouter.PUT("/:socialMediaID", controllers.UpdateSocialMediaRequest)
        socmedRouter.DELETE("/:socialMediaID", controllers.DeleteSocialMedia)
    }
}

func addUserRoutes(r *gin.Engine) {
    userRouter := r.Group("/users")
    {
        userRouter.POST("/register", controllers.UserRegister)
        userRouter.POST("/login", controllers.UserLogin)
        userRouter.PUT("/:userID", controllers.UpdateUser)
        userRouter.DELETE("/:userID", controllers.DeleteUser)
    }
}
