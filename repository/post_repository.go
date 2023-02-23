package repository

import (
	"basic-gin/entity"
	"basic-gin/model"

	"gorm.io/gorm"
)

type PostRepository struct{
	db         *gorm.DB
}
func NewPostRepository(db *gorm.DB) PostRepository{
	return PostRepository{db}
}


func (r *PostRepository) CreatePost(post *entity.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetPostByID( id uint) (entity.Post, error) {
	post := entity.Post{}

	err := r.db.Preload("Comments").First(&post, id).Error
	
	return post, err
}

func (r *PostRepository) GetAllPost() ([]entity.Post, error) {
	var posts[] entity.Post

	err := r.db.Find(&posts).Error

	return posts, err
}

func (r *PostRepository) UpdatePost(ID uint, updatePost *model.UpdatePostRequest) error {
	var post entity.Post

	err := r.db.Model(&post).Where("id = ?", ID).Updates(updatePost).Error

	return err
}

func (r *PostRepository) DeletePost(ID uint) error {
	var post entity.Post

	err := r.db.Delete(&post, ID).Error

	return err
}