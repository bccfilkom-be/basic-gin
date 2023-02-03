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

type commentHandler struct {
	DB *gorm.DB
	Repository repository.CommentRepository
}
// "Constructor" for postHandler
func NewCommentHandler(db *gorm.DB) commentHandler{
	return commentHandler{db, repository.CommentRepository{}}
}

func (h *commentHandler) CreateNewComment(c *gin.Context) {
	// bind incoming http request
	var requestComment model.CreateNewComment
	if err := c.ShouldBindJSON(&requestComment); err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "format body invalid ", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	// create comment
	newComment := entity.Comment{
		Comment: requestComment.Comment,
		PostID: requestComment.PostID,
	}
	err := h.Repository.CreateComment(h.DB, &newComment)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, response.FailOrError(code, "create comment failed", nil))
		return
	}

	//success response
	c.JSON(http.StatusCreated, response.Success(
		"create comment succeeded", newComment,
	))
}

func (h *commentHandler) GetCommentByID(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "invalid id params", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	comment, err := h.Repository.GetCommentByID(h.DB, uint(parsedID))

	if err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(code, "comment not found", nil))
		return
	}

	c.JSON(http.StatusCreated, response.Success(
		"comment found", comment,
	))
}