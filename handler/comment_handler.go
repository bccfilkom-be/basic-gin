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
		code := http.StatusUnprocessableEntity
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

	c.JSON(http.StatusOK, response.Success(
		"comment found", comment,
	))
}

func (h *commentHandler) GetCommentByTitleQuery(c *gin.Context) {
	query := c.Query("comment")

	comments, err := h.Repository.GetCommentByTitleQuery(h.DB, query)

	if err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(code, "comment not found", nil))
		return
	}

	c.JSON(http.StatusOK, response.Success(
		"comment found", comments,
	))
}

func (h *commentHandler) UpdateCommentByID(c *gin.Context) {
	ID := c.Param("id")

	parsedID, err := strconv.ParseUint(ID, 10, 64)

	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "invalid id params", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	var request entity.Comment

	if err := c.ShouldBindJSON(&request); err != nil {
		code := http.StatusUnprocessableEntity
		c.JSON(code, response.FailOrError(
			code, "body is invalid ..", gin.H{
				"error" : err.Error(),
			},
		)) 
		return
	}

	err = h.Repository.UpdateCommentByID(h.DB, uint(parsedID), &request)

	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, response.FailOrError(code, "failed to update comment", nil))
		return
	}

	comment, err := h.Repository.GetCommentByID(h.DB, uint(parsedID))

	if err != nil {
		code := http.StatusNotFound
		c.JSON(code, response.FailOrError(code, "comment not found", nil))
		return
	}

	c.JSON(http.StatusOK, response.Success(
		"comment updated", comment,
	))
}

func (h *commentHandler) DeleteCommentByID(c *gin.Context) {
	ID := c.Param("id")

	parsedID, err := strconv.ParseUint(ID, 10, 64)

	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, response.FailOrError(
			code, "invalid id params", gin.H{
				"error": err.Error(),
			},
		))
		return
	}

	err = h.Repository.DeleteCommentByID(h.DB, uint(parsedID))

	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, response.FailOrError(code, "failed to delete comment", nil))
		return
	}

	c.JSON(http.StatusOK, response.Success(
		"comment deleted", nil,
	))
}
