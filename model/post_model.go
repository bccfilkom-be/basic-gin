package model

type CreatePostRequest struct {
	Title   string `binding:"required"`
	Content string `binding:"required"`
}