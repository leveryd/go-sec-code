package handler

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	_, err := c.Writer.Write([]byte("pong"))
	if err != nil {
		return
	}
}
