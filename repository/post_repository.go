package repository

import (
	"basic-gin/entity"

	"gorm.io/gorm"
)

type PostRepository struct{}

func (r *PostRepository) CreatePost(db *gorm.DB, post *entity.Post) error {
	return db.Create(post).Error
}