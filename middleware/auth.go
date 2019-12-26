package middleware

import (
	"blog_go/pkg"
	"blog_go/util/e"
	"github.com/gin-gonic/gin"
)

func Verification() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			e.AbortJson(c, &e.Return{Code:e.TOKEN_IN_VAIN})
			return
		}
		user, err := pkg.ParseToken(token)
		if err != nil {
			e.AbortJson(c, &e.Return{Code:e.TOKEN_IN_VAIN})
			return
		}
		c.Set("login_user", user)

		c.Next()
	}
}
