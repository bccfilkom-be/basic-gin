package handler

import (
	"basic-gin/entity"
	"basic-gin/model"
	"basic-gin/repository"
	"basic-gin/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postHandler struct {
	DB         *gorm.DB
	Repository repository.PostRepository
}

// "Constructor" for postHandler
func NewPostHandler(db *gorm.DB) postHandler {
	return postHandler{db, repository.PostRepository{}}
}

func (h *postHandler) CreatePost(c *gin.Context) {
	// bind incoming http request
	request := model.CreatePostRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		code := http.StatusUnprocessableEntity
		c.JSON(code, response.FailOrError(
			code, "Create post failed", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	// create post
	post := entity.Post{
		Title:   request.Title,
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

func (h *postHandler) GetPostByID(c *gin.Context) {
	// binding param to request model
	request := model.GetPostByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "Get post failed", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	// find post
	post, err := h.Repository.GetPostByID(h.DB, request.ID)
	if err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(
			code, "Post not found", gin.H{
				"error": "404 not found",
			},
		))
		return
	}

	//success
	c.JSON(http.StatusOK, response.Success("Post found", post))
}

func (h *postHandler) GetAllPost(c *gin.Context) {
	posts, err := h.Repository.GetAllPost(h.DB)

	if err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(
			code, "Posts not found", gin.H{
				"error": "404 not found",
			},
		))
		return
	}

	c.JSON(http.StatusOK, response.Success("Posts Found", posts))
}

func (h *postHandler) UpdatePostByID(c *gin.Context) {
	ID := c.Param("id")

	var request model.UpdatePostRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "body is invalid ..", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	parsedID, _ := strconv.ParseUint(ID, 10, 64)

	request = model.UpdatePostRequest{
		Title:   request.Title,
		Content: request.Content,
	}

	err := h.Repository.UpdatePost(h.DB, uint(parsedID), &request)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, response.FailOrError(code, "Update post failed", nil))
		return
	}

	post, err := h.Repository.GetPostByID(h.DB, uint(parsedID))
	if err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(
			code, "Post not found", gin.H{
				"error": "404 not found",
			},
		))
		return
	}

	//success
	c.JSON(http.StatusOK, response.Success("updated post successfully", post))
}

func (h *postHandler) DeletePostByID(c *gin.Context) {
	ID := c.Param("id")

	parsedID, _ := strconv.ParseUint(ID, 10, 64)

	err := h.Repository.DeletePost(h.DB, uint(parsedID))

	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, response.FailOrError(code, "delete post failed", nil))
		return
	}

	c.JSON(http.StatusOK, response.Success("successfully deleted post", nil))
}
