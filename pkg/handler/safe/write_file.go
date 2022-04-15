package safe

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)

func GoodUploadFile(c *gin.Context) {
	c.Request.ParseMultipartForm(32 << 20)
	file, handler, err := c.Request.FormFile("uploadfile")
	if err != nil {
		c.String(400, "Bad request")
		return
	}
	defer file.Close()

	// safe check
	if strings.Contains(handler.Filename, "..") {
		c.String(400, "Bad request")
		return
	}

	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.String(400, "Bad request")
		return
	}
	defer f.Close()
	io.Copy(f, file)
	c.String(200, "Upload success")
}
