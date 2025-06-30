package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

type CreateComment struct {
	Content string `json:"content" binding:"required"`
}
