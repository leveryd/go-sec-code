package filter

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	c.Writer.Header().Set("auth", "success")
	c.Next()
}
