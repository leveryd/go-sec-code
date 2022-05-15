package handler

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Ping(c *gin.Context) {
	_, err := c.Writer.Write([]byte("pong"))
	if err != nil {
		return
	}
}

func SimulateLongTask(_ *gin.Context) {
	time.Sleep(10 * time.Second)
	println("long task done")
}
