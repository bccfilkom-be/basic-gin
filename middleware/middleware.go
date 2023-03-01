package middlware

import (
	"basic-gin/model"
	sdk_jwt "basic-gin/sdk/jwt"
	"basic-gin/sdk/response"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
Definisi dari middleware sendiri versi penulis,
sebuah blok kode yang dipanggil sebelum ataupun sesudah http request di proses.

Kita bisa menggunakan middleware buat ngecek Jwt token yang dikirim.
Tujuannya untuk memperbolehkan atau melarang request mengakses endpoint yang private
*/
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Umumnya, Jwt Token dikirim lewat Header Http 'Authorization' dengan nilai Bearer jwt_token
		authorization := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			c.Abort()
			msg := "wrong header value"
			response.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
			return
		}
		tokenJwt := authorization[7:] // menghilangkan Bearer
		claims := model.UserClaims{} // user claims yg sudah didefinisikan dari model
		jwtKey := os.Getenv("secret_key")
		if err := sdk_jwt.DecodeToken(tokenJwt, &claims, jwtKey); err != nil {
			c.Abort()
			response.FailOrError(c, http.StatusUnauthorized, "unauthorized", err)
			return
		}
		c.Set("user", claims)
	}
}
