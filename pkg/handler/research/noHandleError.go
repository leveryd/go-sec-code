package research

import (
	"github.com/gin-gonic/gin"
)

// NoHandleError /risk/error'
func NoHandleError(c *gin.Context) {
	panic("test")
	//c.Writer.Write([]byte(c.ClientIP()))
}
