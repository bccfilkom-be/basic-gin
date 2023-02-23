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

type commentHandler struct {
	Repository repository.CommentRepository
}
// "Constructor" for postHandler
func NewCommentHandler(repo *repository.CommentRepository) commentHandler{
	return commentHandler{*repo}
}

func (h *commentHandler) CreateNewComment(c *gin.Context) {
	// bind incoming http request
	var requestComment model.CreateNewComment
	if err := c.ShouldBindJSON(&requestComment); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "create new comment failed", err)
		return
	}

	// create comment
	newComment := entity.Comment{
		Comment: requestComment.Comment,
		PostID: requestComment.PostID,
	}
	err := h.Repository.CreateComment(&newComment)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create comment failed", err)
		return
	}

	//success response
	response.Success(c, http.StatusCreated, "create comment succeeded", newComment)
}

func (h *commentHandler) GetCommentByID(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}

	comment, err := h.Repository.GetCommentByID(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "comment not found", err)
		return
	}

	response.Success(c, http.StatusOK, "comment found", comment)
}

func (h *commentHandler) GetCommentByTitleQuery(c *gin.Context) {
	query := c.Query("comment")

	comments, err := h.Repository.GetCommentByTitleQuery(query)

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "comment not found", err)
		return
	}

	response.Success(c, http.StatusOK, "comment found", comments)
}

func (h *commentHandler) UpdateCommentByID(c *gin.Context) {
	ID := c.Param("id")

	parsedID, err := strconv.ParseUint(ID, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}

	var request entity.Comment

	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "body is invalid ..", err) 
		return
	}

	err = h.Repository.UpdateCommentByID(uint(parsedID), &request)

	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "failed to update comment", err)
		return
	}

	comment, err := h.Repository.GetCommentByID(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "comment not found", err)
		return
	}

	response.Success(c, http.StatusOK, "comment updated", comment)
}

func (h *commentHandler) DeleteCommentByID(c *gin.Context) {
	ID := c.Param("id")

	parsedID, err := strconv.ParseUint(ID, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}

	err = h.Repository.DeleteCommentByID(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "failed to delete comment", err)
		return
	}

	response.Success(c, http.StatusOK, "comment deleted", err)
}
