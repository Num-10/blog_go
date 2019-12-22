package middleware

import "github.com/gin-gonic/gin"

func Verification(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.AbortWithStatusJSON(200, gin.H{
			"message": "no token",
		})
	}
}
