package handler

import (
	"basic-gin/entity"
	"basic-gin/model"
	"basic-gin/repository"
	"basic-gin/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	Repository repository.PostRepository
}

// "Constructor" for postHandler
func NewPostHandler(repo *repository.PostRepository) postHandler {
	return postHandler{*repo}
}

func (h *postHandler) CreatePost(c *gin.Context) {
	claimsTemp, _ := c.Get("user")
	claims := claimsTemp.(model.UserClaims)

	// bind incoming http request
	request := model.CreatePostRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Create post failed", err)
		return
	}

	// create post
	post := entity.Post{
		UserID: claims.ID,
		Title:   request.Title,
		Content: request.Content,
	}
	err := h.Repository.CreatePost(&post)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create post failed", err)
		return
	}

	//success response
	response.Success(c, http.StatusCreated, "Post creation succeeded",request)
}

func (h *postHandler) GetPostByID(c *gin.Context) {
	// binding param to request model
	request := model.GetPostByIDRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Get post failed", err)
		return
	}

	// find post
	post, err := h.Repository.GetPostByID(request.ID)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Post not found", err)
		return
	}

	//success
	response.Success(c, http.StatusOK, "Post found", post)
}

func (h *postHandler) GetAllPost(c *gin.Context) {
	posts, err := h.Repository.GetAllPost()

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Posts not found", err)
		return
	}

	response.Success(c, http.StatusOK, "Posts Found", posts)
}

func (h *postHandler) UpdatePostByID(c *gin.Context) {
	ID := c.Param("id")

	var request model.UpdatePostRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "body is invalid ..", err)
		return
	}

	parsedID, _ := strconv.ParseUint(ID, 10, 64)

	request = model.UpdatePostRequest{
		Title:   request.Title,
		Content: request.Content,
	}

	err := h.Repository.UpdatePost(uint(parsedID), &request)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Update post failed", err)
		return
	}

	post, err := h.Repository.GetPostByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Post not found", err)
		return
	}

	//success
	response.Success(c, http.StatusOK, "updated post successfully", post)
}

func (h *postHandler) DeletePostByID(c *gin.Context) {
	ID := c.Param("id")

	parsedID, _ := strconv.ParseUint(ID, 10, 64)

	err := h.Repository.DeletePost(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "delete post failed", err)
		return
	}

	response.Success(c, http.StatusOK, "successfully deleted post", nil)
}
