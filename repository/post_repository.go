package repository

import (
	"basic-gin/entity"
	"basic-gin/model"

	"gorm.io/gorm"
)

type PostRepository struct{}

func (r *PostRepository) CreatePost(db *gorm.DB, post *entity.Post) error {
	return db.Create(post).Error
}

func (r *PostRepository) GetPostByID(db *gorm.DB, id uint) (entity.Post, error) {
	post := entity.Post{}

	err := db.Preload("Comments").First(&post, id).Error
	
	return post, err
}

func (r *PostRepository) GetAllPost(db *gorm.DB) ([]entity.Post, error) {
	var posts[] entity.Post

	err := db.Find(&posts).Error

	return posts, err
}

func (r *PostRepository) UpdatePost(db *gorm.DB, ID uint, updatePost *model.UpdatePostRequest) error {
	var post entity.Post

	err := db.Model(&post).Where("id = ?", ID).Updates(updatePost).Error

	return err
}

func (r *PostRepository) DeletePost(db *gorm.DB, ID uint) error {
	var post entity.Post

	err := db.Delete(&post, ID).Error

	return err
}