package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    CommentRoutes(router)
    PhotosRoutes(router)
    SocmedRoutes(router)
    UserRoutes(router)

    return router
}

func CommentRoutes(r *gin.Engine) {
    commentRouter := r.Group("/comments")
    {
        commentRouter.Use(middlewares.Authentication())
        commentRouter.POST("/", controllers.CreateComment)
        commentRouter.GET("/", controllers.GetComment)
        commentRouter.GET("/:commentID", controllers.GetCommentByID)
        commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateCommentRequest)
        commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
    }
}

func PhotosRoutes(r *gin.Engine) {
    photosRouter := r.Group("/photos")
    {
        photosRouter.Use(middlewares.Authentication())
        photosRouter.POST("/", controllers.CreatePhoto)
        photosRouter.GET("/", controllers.GetPhotos)
        photosRouter.GET("/:photoID", controllers.GetPhotosByID)
        photosRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
        photosRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
    }
}

func SocmedRoutes(r *gin.Engine) {
    socmedRouter := r.Group("/socialmedias")
    {
        socmedRouter.Use(middlewares.Authentication())
        socmedRouter.POST("/", controllers.CreateSocialMediaRequest)
        socmedRouter.GET("/", controllers.GetSocialMedia)
        socmedRouter.GET("/:socialMediaID", controllers.GetSocialMediaByID)
        socmedRouter.PUT("/:socialMediaID", middlewares.SocmedAuthorization(), controllers.UpdateSocialMediaRequest)
        socmedRouter.DELETE("/:socialMediaID", middlewares.SocmedAuthorization(), controllers.DeleteSocialMedia)
    }
}

func UserRoutes(r *gin.Engine) {
    userRouter := r.Group("/users")
    {
        userRouter.POST("/register", controllers.UserRegister)
        userRouter.POST("/login", controllers.UserLogin)
        userRouter.PUT("/:userID", middlewares.UserAuthorization(), controllers.UpdateUser)
        userRouter.DELETE("/:userID", middlewares.UserAuthorization(), controllers.DeleteUser)
    }
}
