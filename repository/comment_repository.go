package repository

import (
	"basic-gin/entity"
	"errors"

	"gorm.io/gorm"
)

type CommentRepository struct{
	db         *gorm.DB
}
func NewCommentRepository(db *gorm.DB) CommentRepository{
	return CommentRepository{db}
}

func (r *CommentRepository) CreateComment(comment *entity.Comment) error {
	res := r.db.Create(comment)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("no rows affected, failed to create comment")
	}

	return nil
}

func (r *CommentRepository) GetCommentByID(ID uint) (*entity.Comment, error) {
	var comment entity.Comment

	result := r.db.First(&comment, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func (r *CommentRepository) GetCommentByTitleQuery(comm string) (*[]entity.Comment, error) {
	var comment[] entity.Comment

	search := "%" + comm + "%"

	result := r.db.Model(&comment).Where("comment like ?", search).Find(&comment)

	if result.Error != nil {
		return nil, result.Error
	}

	// ini bebas mau digunakan atau tidak..
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected, posts not found")
	}

	return &comment, nil
}

func (r *CommentRepository) UpdateCommentByID( ID uint, updateComment *entity.Comment) error {
	var comment entity.Comment

	result := r.db.Model(&comment).Where("ID = ?", ID).Updates(updateComment)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows affected, failed to update comment")
	}

	return nil
}

func (r *CommentRepository) DeleteCommentByID( ID uint) error {
	var comment entity.Comment

	result := r.db.Delete(&comment, ID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows affected, missing comment or abnormal behaviour happen")
	}
	return nil
}