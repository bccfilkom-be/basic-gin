package handler

import (
	"basic-gin/entity"
	"basic-gin/model"
	"basic-gin/repository"
	"basic-gin/sdk/crypto"
	sdk_jwt "basic-gin/sdk/jwt"
	"basic-gin/sdk/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	Repository repository.UserRepository
}
func NewUserHandler(repo *repository.UserRepository) userHandler {
	return userHandler{*repo}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var request model.RegisterUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	// Ingat, sebelum menyimpan data user ke database, sebaiknya lakukan hashing password terlebih dahulu
	hashedPassword, err := crypto.HashValue(request.Password)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "user creation failed", err)
		return
	}

	user := entity.User{
		Name: request.Name,
		Username: request.Username,
		Password: hashedPassword,
	}

	err = h.Repository.CreateUser(&user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create user failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "success create user", user)
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
	// `dto` seharusnya di model yo, tapi ini contoh doang
	req := struct {
		ID uint `uri:"id" binding:"required"`
	}{}
	err := c.ShouldBindUri(&req)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad uri", err)
		return
	}
	result, err := h.Repository.GetUserById(req.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailOrError(c, http.StatusNotFound, "user not found", err)
			return
		}
		response.FailOrError(c, http.StatusInternalServerError, "get user failed", err)
		return
	}
	response.Success(c, http.StatusOK, "success get user", result)
}
