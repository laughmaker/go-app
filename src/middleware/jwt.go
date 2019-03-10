package middleware

import (
	"net/http"

	"app/src/pkg/app"
	"app/src/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = app.SUCCESS
		token := c.Query("token")

		if token == "" {
			code = app.TOKEN_EMPTY
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = app.TOKEN_EXPIRED
				default:
					code = app.INVALID_TOKEN
				}
			}
		}

		resp := app.Resp{C: c}
		if code != app.SUCCESS {
			resp.Error(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
