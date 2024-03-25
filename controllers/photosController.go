package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// var (
// 	appJSON = "application/json"
// )



func CreatePhoto(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": Photo.ID,
		"caption":  Photo.Caption,
		"title":    Photo.Title,
		"photo_url": Photo.PhotoUrl,
		"user_id":  Photo.UserID,
	})
}

func GetPhotos(c *gin.Context) {
    db := database.GetDB()
    contentType := helpers.GetContentType(c)

	
    var photos []models.Photo
	if contentType == appJSON {
		c.ShouldBindJSON(&photos)
	} else {
		c.ShouldBind(&photos)
	}
    err := db.Debug().Find(&photos).Error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    var responseData []gin.H
    for _, photo := range photos {
        user := models.User{}            
        db.First(&user, photo.UserID) 

        responseData = append(responseData, gin.H{
            "id":        photo.ID,
            "caption":   photo.Caption,
            "title":     photo.Title,
            "photo_url": photo.PhotoUrl,
            "user_id":   photo.UserID,
            "user": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
        })
    }

    c.JSON(http.StatusOK, responseData)
}


func GetPhotosByID(c *gin.Context) {
    db := database.GetDB()
    contentType := helpers.GetContentType(c)

    var photo []models.Photo

	if contentType == appJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

    err := db.Debug().Where("id = ?", c.Param("photoID")).First(&photo).Error

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    var responseData []gin.H
    for _, photo := range photo {
        user := models.User{}            
        db.First(&user, photo.UserID) 

        responseData = append(responseData, gin.H{
            "id":        photo.ID,
            "caption":   photo.Caption,
            "title":     photo.Title,
            "photo_url": photo.PhotoUrl,
            "user_id":   photo.UserID,
            "user": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
        })
    }

    c.JSON(http.StatusOK, responseData)
}


func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Caption: Photo.Caption, Title: Photo.Title, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       Photo.ID,
		"caption":  Photo.Caption,
		"title":    Photo.Title,
		"photo_url": Photo.PhotoUrl,
		"user_id":  Photo.UserID,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Debug().Where("id = ?", c.Param("photoID")).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
	})
}