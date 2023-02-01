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
			code, "Create post failed",
			map[string]interface{}{
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