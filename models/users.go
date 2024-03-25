package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"size:50;not null"`
	Email        string        `gorm:"size:150;not null"`
	Password     string        `gorm:"type:text;not null"`
	Age          int           `gorm:"not null"`
	Photos       []Photo       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Comments     []Comment     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,gte=8"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// type UpdateCurrentUserInput struct {
// 	Username string `json:"username" binding:"required"`
// 	Email string `json:"email" binding:"required"`
// 	Age int `json:"age" binding:"required"`
// 	ProfileImageURL string `json:"profile_image_url,omitempty" validate:"omitempty"`
// }

func (u *RegisterInput) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}