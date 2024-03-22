package models

import (
	"time"
)

type SocialMedia struct {
	ID int64 `gorm:"primary_key"`
	Name string `gorm:"size:50";"not null"`
	SocialMediaUrl string `gorm:"type:text";"not null"`
	UserID int64 `gorm:"not null"`
	User User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateSocialMediaRequest struct {
	Name string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type UpdateSocialMediaRequest struct {
	Name string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}