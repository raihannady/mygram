package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	// "regexp"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
    db := database.GetDB()
    contentType := helpers.GetContentType(c)
    User := models.User{}

    
    if contentType == appJSON {
        c.ShouldBindJSON(&User)
    } else {
        c.ShouldBind(&User)
    }

    // var registerInput models.RegisterInput

    // if contentType == appJSON {
    //     log.Println("JSON request received:", registerInput)
    //     c.ShouldBindJSON(&registerInput)
    //     log.Println("Parsed User:", registerInput)
    // } else {
    //     c.ShouldBind(&registerInput)
    // }
    
    // if err := c.ShouldBindJSON(&registerInput); err != nil {
    //     c.JSON(http.StatusBadRequest, gin.H{
    //         "error":   "Bad Request",
    //         "message": err.Error(),
    //     })
    //     return
    // }

    // var existingUser models.User
    // if err := db.Where("email = ?", registerInput.Email).First(&existingUser).Error; err == nil {
    //     c.JSON(http.StatusConflict, gin.H{
    //         "error":   "Conflict",
    //         "message": "Email already exists",
    //     })
    //     return
    // }

	// var existingUsername models.User
	// if err := db.Where("username = ?", registerInput.Username).First(&existingUsername).Error; err == nil {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"error":   "Conflict",
	// 		"message": "Username already exists",
	// 	})
	// 	return
	// }

    // User.Email = registerInput.Email
    // User.Username = registerInput.Username
    // User.Age = registerInput.Age
    User.Password = helpers.HashPassword(User.Password)


    err := db.Debug().Create(&User).Error

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "id":       User.ID,
        "email":    User.Email,
        "username": User.Username,
        "age":      User.Age,
        "profile_image_url": User.ProfileImage,
    })
}


func UserLogin(c *gin.Context) {
    db := database.GetDB()
    contentType := helpers.GetContentType(c)
    _, _ = db, contentType
    User := models.User{}
    password := ""

    var loginInput models.SignInInput
    if err := c.ShouldBindJSON(&loginInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(), 
        })
        return
    }

    User.Email = loginInput.Email
    password = loginInput.Password

    err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error":   "Unauthorized",
            "message": "Invalid email/password",
        })
        return
    }

    comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

    if !comparePass {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error":   "Unauthorized",
            "message": "Invalid email/password",
        }) 
        return
    }

    token := helpers.GenerateToken(User.ID, User.Email)

    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}


func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Where("id = ?", c.Param("userID")).Updates(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": User.ID,
		"email":  User.Email,
		"username":   User.Username,
		"age":    User.Age,
        "profile_image_url": User.ProfileImage,
	})
	
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Where("id = ?", c.Param("userID")).Delete(&User).Error

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