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

func CreateSocialMediaRequest(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)

}

func GetSocialMedia(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var socialmedia []models.SocialMedia

	if contentType == appJSON {
		c.ShouldBindJSON(&socialmedia)
	} else {
		c.ShouldBind(&socialmedia)
	}

	var err error

	err = db.Debug().Where("user_id = ?", userID).Find(&socialmedia).Error

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }


	var responseData []gin.H
    for _, socialmedia := range socialmedia {
        user := models.User{}            
        db.First(&user, socialmedia.UserID)

        responseData = append(responseData, gin.H{
            "id":        socialmedia.ID,
            "name":      socialmedia.Name,
            "user_id":   socialmedia.UserID,
            "user": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
        })
    }

    c.JSON(http.StatusOK, responseData)
}

func GetSocialMediaByID(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	var socialmedia []models.SocialMedia

	err := db.Debug().Where("id = ?", c.Param("socialMediaID")).First(&socialmedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var responseData []gin.H
    for _, socialmedia := range socialmedia {
        user := models.User{}            
        db.First(&user, socialmedia.UserID)

        responseData = append(responseData, gin.H{
            "id":        socialmedia.ID,
            "name":      socialmedia.Name,
            "user_id":   socialmedia.UserID,
            "user": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
        })
    }

    c.JSON(http.StatusOK, responseData)
}

func UpdateSocialMediaRequest(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socmedId, _ := strconv.Atoi(c.Param("socialMediaID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socmedId)


	err := db.Model(&SocialMedia).Where("id = ?", socmedId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Where("id = ?", c.Param("socialMediaID")).Delete(&SocialMedia).Error

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