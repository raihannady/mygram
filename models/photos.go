package models

type Photo struct {
	GormModel
	Title    string `gorm:"size:100;not null"`
	Caption  string `gorm:"size:200"`
	PhotoUrl string `gorm:"type:text;not null"`
	UserID   uint
	User     *User
	Comments *Comment
}

type CreatePhotoRequest struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption,omitempty"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type UpdatePhoto struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption,omitempty"`
	PhotoURL string `json:"photo_url" validate:"required"`
}