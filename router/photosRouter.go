package router

import (
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func PhotosRouter(router *gin.RouterGroup)  {

    photosRouter := router.Group("/photos")
    {
        photosRouter.POST("/", controllers.CreatePhoto)
        photosRouter.GET("/", controllers.GetPhotos)
        photosRouter.GET("/:photoID", controllers.GetPhotosByID)
        photosRouter.PUT("/:photoID", controllers.UpdatePhoto)
        photosRouter.DELETE("/:photoID", controllers.DeletePhoto)
    }
}
