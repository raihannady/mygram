package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"size:50;not null"`
	SocialMediaUrl string `gorm:"type:text;not null"`
	UserID         uint
	User           *User
}

type CreateSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"url" binding:"required"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"url" validate:"required"`
}