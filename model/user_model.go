package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type RegisterUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserClaims struct {
	ID uint `json:"id"`// bebas dahh mau pake 'Subject' di registered claims juga bebas 
	jwt.RegisteredClaims
}
// exp example: (time.Hour * 1), etc
func NewUserClaims(id uint, exp time.Duration) UserClaims{
	return UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}