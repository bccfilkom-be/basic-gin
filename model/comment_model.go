package model

type CreateNewComment struct {
	Comment string `json:"comment" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}