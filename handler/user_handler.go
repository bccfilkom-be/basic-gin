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
	db         *gorm.DB
	Repository repository.UserRepository
}

func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		db:         db,
		Repository: repository.UserRepository{},
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user model.RegisterUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(http.StatusBadRequest, err.Error(), nil)
		return
	}
	result, err := h.Repository.CreateUser(h.db, user)
	if err != nil {
		response.FailOrError(http.StatusInternalServerError, err.Error(), nil)
		return
	}
	c.JSON(http.StatusCreated, response.Success("success create user", result))
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var user model.LoginUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(http.StatusBadRequest, err.Error(), nil)
		return
	}
	result, err := h.Repository.LoginUser(h.db, user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.FailOrError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *userHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	result, err := h.Repository.GetUserById(h.db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailOrError(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.Success("success get user", result))
}
