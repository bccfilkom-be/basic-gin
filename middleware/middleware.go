package middlware

import (
	"basic-gin/sdk/response"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
		// validate jwt adalah token yang sudah divalidasi
		validateJwt, err := jwt.Parse(tokenJwt, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("secret_key")), nil
		})
		if err != nil {
			c.Abort()
			response.FailOrError(c, http.StatusForbidden, err.Error(), err)
			return
		}
		// jwtFix adalah bentuk asli token nya
		jwtFix, ok := validateJwt.Claims.(jwt.MapClaims)
		if !ok {
			c.Abort()
			response.FailOrError(c, http.StatusForbidden, "data token jwt tidak valid", nil)
			return
		}
		// Token tidak valid
		if jwtFix.Valid() != nil {
			c.Abort()
			response.FailOrError(c, http.StatusForbidden, jwtFix.Valid().Error(), jwtFix.Valid())
			return
		} else {
			// Token valid
			c.Set("user", jwtFix)
			c.Next()
		}
	}
}
