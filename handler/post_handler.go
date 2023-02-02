package handler

import (
	"basic-gin/entity"
	"basic-gin/model"
	"basic-gin/repository"
	"basic-gin/sdk/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postHandler struct {
	DB *gorm.DB
	Repository repository.PostRepository
}
// "Constructor" for postHandler
func NewPostHandler(db *gorm.DB) postHandler{
	return postHandler{db, repository.PostRepository{}}
}

func (h *postHandler) CreatePost(c *gin.Context){
	// bind incoming http request
	request := model.CreatePostRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "Create post failed", gin.H{
				"error" : err.Error(),
			},
		)) 
		return
	}

	// create post
	post := entity.Post{
		Title: request.Title,
		Content: request.Content,
	}
	err := h.Repository.CreatePost(h.DB, &post)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, response.FailOrError(code, "Create post failed", nil))
		return
	}

	//success response
	c.JSON(http.StatusCreated, response.Success(
		"Post creation succeeded",
		 request,
	))
}

func (h *postHandler) GetPostByID(c *gin.Context){
	// binding param to request model
	request := model.GetPostByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "Get post failed", gin.H{
				"error" : err.Error(),
			},
		))
		return
	}

	// find post
	post, err := h.Repository.GetPostByID(h.DB, request.ID)
	if  err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(
			code, "Post not found", gin.H{
				"error" : "404 not found",
			},
		))
		return
	}

	//success
	c.JSON(http.StatusOK, response.Success("Post found", post))
}