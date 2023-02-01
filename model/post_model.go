package model

type CreatePostRequest struct {
	Title   string `binding:"required"`
	Content string `binding:"required"`
}

type GetPostByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}