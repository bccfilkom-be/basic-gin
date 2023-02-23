package handler

import (
	"basic-gin/model"
	"basic-gin/repository"
	"basic-gin/sdk/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	Repository repository.UserRepository
}

func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		Repository: repository.UserRepository{},
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user model.RegisterUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}
	result, err := h.Repository.CreateUser(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create user failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "success create user", result)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var user model.LoginUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}
	result, err := h.Repository.LoginUser(user)
	if err != nil {
		//TODO: ^INI NGAWUR repository kok kyk service 
		response.FailOrError(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *userHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	result, err := h.Repository.GetUserById(id)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "get user failed", err)
		return
	}
	response.Success(c, http.StatusOK, "success get user", result)
}
