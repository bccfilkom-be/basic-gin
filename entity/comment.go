package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Comment string `json:"comment"`
	// UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}