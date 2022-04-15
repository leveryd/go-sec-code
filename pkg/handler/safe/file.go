package safe

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

// FileRead safe/fileread?fpath=../../../../../../etc/hosts
func FileRead(c *gin.Context) {
	filePath := c.Query("fpath")
	filePath = filepath.Clean("/" + filePath) // safe
	filePath = filepath.Join("./", filePath)

	openFile, err := os.Open(filePath)
	if err != nil {
		c.JSON(200, "read file fail")
		return
	}
	b := make([]byte, 100)
	openFile.Read(b)
	c.Writer.Write(b)
}
