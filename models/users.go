package models

import (
	"mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID int64 `gorm:"primary_key"`
	Username string `gorm:"size:50";"not null"`
	Email string `gorm:"size:150";"not null"`
	Password string `gorm:"type:text";"not null"`
	Age int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos []Photo `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Comments []Comment `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	SocialMedias []SocialMedia `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`	
}

type SignInInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateCurrentUserInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age int `json:"age" binding:"required"`
	ProfileImageURL string `json:"profile_image_url,omitempty" validate:"omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}