package safe

import (
	"github.com/gin-gonic/gin"
	"os"
)

func GoodUploadFile(c *gin.Context) {
	file, err := c.FormFile("filename")
	if err != nil {
		return
	}
	if err != nil {
		c.String(400, "Bad request")
		return
	}

	// safe check
	//if strings.Contains(file.Filename, "..") {
	//	c.String(400, "Bad request")
	//	return
	//}

	os.Mkdir("./upload", os.ModePerm)
	filepath := "./upload/" + file.Filename

	{
		err := c.SaveUploadedFile(file, filepath)
		if err != nil {
			c.String(200, "Upload failed")
			return
		}
	}

	defer os.Remove(filepath)

	c.String(200, "Upload success")
}

/*
FormFile函数会保证文件名不包含.. ，所以不需要判断文件名中是否有 ..
*/
