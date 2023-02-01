package repository

import (
	"basic-gin/entity"

	"gorm.io/gorm"
)

type PostRepository struct{}

func (r *PostRepository) CreatePost(db *gorm.DB, post *entity.Post) error {
	return db.Create(post).Error
}

func (r *PostRepository) GetPostByID(db *gorm.DB, id uint) (entity.Post, error) {
	post := entity.Post{}
	err := db.First(&post, id).Error
	return post, err
}