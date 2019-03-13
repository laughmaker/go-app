package middleware

import (
	"app/src/pkg/log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获panic异常
		defer log.Try(c)

		c.Next()
	}
}
