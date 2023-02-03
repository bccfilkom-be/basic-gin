package repository

import (
	"basic-gin/entity"
	"errors"

	"gorm.io/gorm"
)

type CommentRepository struct{}

func (r *CommentRepository) CreateComment(db *gorm.DB, comment *entity.Comment) error {
	res := db.Create(comment)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("no rows affected, failed to create comment")
	}

	return nil
}

func (r *CommentRepository) GetCommentByID(db *gorm.DB, ID uint) (*entity.Comment, error) {
	var comment entity.Comment

	result := db.First(&comment, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func (r *CommentRepository) GetCommentByTitleQuery(db *gorm.DB, comm string) (*[]entity.Comment, error) {
	var comment[] entity.Comment

	search := "%" + comm + "%"

	result := db.Model(&comment).Where("comment like ?", search).Find(&comment)

	if result.Error != nil {
		return nil, result.Error
	}

	// ini bebas mau digunakan atau tidak..
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected, posts not found")
	}

	return &comment, nil
}

