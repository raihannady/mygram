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

func CreateComment(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	var CreateComment models.CreateComment
	Comment.PhotoID = CreateComment.PhotoID
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": Comment.ID,
		"message": Comment.Message,
		"photo_id": Comment.PhotoID,
		"user_id": Comment.UserID,
	})

}




func GetComment(c *gin.Context) {
	
	db := database.GetDB()
    contentType := helpers.GetContentType(c)

	
    var comment []models.Comment
	if contentType == appJSON {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

    err := db.Debug().Find(&comment).Error
	
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    var responseData []gin.H
    for _, comment := range comment {
        user := models.User{}            
        db.First(&user, comment.UserID) 
		photo := models.Photo{}
		db.First(&photo, comment.PhotoID)

        responseData = append(responseData, gin.H{
            "id":        comment.ID,
            "caption":   comment.Message,
            "photo_id":  comment.PhotoID,
            "user_id":   comment.UserID,
            "user": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
			"photo": gin.H{
				"id": photo.ID,
				"title": photo.Title,
				"caption": photo.Caption,
				"photo_url": photo.PhotoUrl,
				"user_id": photo.UserID,
			},
        })
    }

    c.JSON(http.StatusOK, responseData)
}

func GetCommentByID(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	var comment []models.Comment


	err := db.Debug().Where("id = ?", c.Param("commentID")).First(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var responseData []gin.H
    for _, comment := range comment {
        user := models.User{}            
        db.First(&user, comment.UserID) 
		photo := models.Photo{}
		db.First(&photo, comment.PhotoID)

        responseData = append(responseData, gin.H{
            "id":        comment.ID,
            "caption":   comment.Message,
            "photo_id":  comment.PhotoID,
            "user_id":   comment.UserID,
            "user": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
			"photo": gin.H{
				"id": photo.ID,
				"title": photo.Title,
				"caption": photo.Caption,
				"photo_url": photo.PhotoUrl,
				"user_id": photo.UserID,
			},
        })
    }

    c.JSON(http.StatusOK, responseData)
}

func UpdateCommentRequest(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message , PhotoID: Comment.PhotoID}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      Comment.ID,
		"message": Comment.Message,
		"photo_id": Comment.PhotoID,
		"user_id": Comment.UserID,
	})
}

func DeleteComment(c *gin.Context) {
    db := database.GetDB()
    contentType := helpers.GetContentType(c)
    Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

    err := db.Debug().Where("id = ?", c.Param("commentID")).Delete(&Comment).Error

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Reques",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
    })
}
