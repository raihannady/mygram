package models

type Comment struct {
	GormModel
	Message string `gorm:"size:200;not null"`
	UserID  uint
	User    *User
	PhotoID uint
	Photo   *Photo
}

type CreateComment struct {
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" validate:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required"`
}

// func (u *CreateComment) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(u)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}
// 	return
// }