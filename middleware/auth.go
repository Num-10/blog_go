package middleware

import (
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
)

func Verification() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			e.AbortJson(c, &e.Return{Code:e.TOKEN_IN_VAIN})
		}
	}
}
