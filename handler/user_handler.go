package handler

import (
	"basic-gin/model"
	"basic-gin/repository"
	"basic-gin/sdk/crypto"
	sdk_jwt "basic-gin/sdk/jwt"
	"basic-gin/sdk/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Repository repository.UserRepository
}
func NewUserHandler(repo *repository.UserRepository) userHandler {
	return userHandler{*repo}
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
	var request model.LoginUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	// get email
	user, err := h.Repository.FindByUsername(request.Username)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "email not found", err)
		return
	}

	// compare password
	err = crypto.ValidateHash(request.Password, user.Password)
	if err != nil {
		msg := "wrong password"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}
	// create jwt
	tokenJwt, err := sdk_jwt.GenerateToken(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}

	// success response
	response.Success(c, http.StatusOK, "login success", gin.H{
		"token" : tokenJwt,
	})
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
