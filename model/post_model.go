package model

type CreatePostRequest struct {
	Title   string `binding:"required" json:"title"`
	Content string `binding:"required" json:"content"`
}

type GetPostByIDRequest struct {
	ID uint `uri:"id" binding:"required"`
}

type UpdatePostRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
}