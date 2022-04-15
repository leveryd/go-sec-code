package unsafe

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
)

func BadFileRead1(c *gin.Context) {

	filePath := c.Query("fpath")
	filePath = filepath.Clean(filePath)
	filePath = filepath.Join("./", filePath)

	openFile, err := os.Open(filePath)
	if err != nil {
		c.JSON(200, "read file fail")
		return
	}
	b := make([]byte, 100)

	{
		_, err := openFile.Read(b)
		if err != nil {
			return
		}
	}

	{
		_, err := c.Writer.Write(b)
		if err != nil {
			return
		}
	}
}

func BadFileRead2(c *gin.Context) {
	filePath := c.Query("fpath")
	filePath = filepath.Clean(filePath)
	filePath = filepath.Join("./", filePath)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.JSON(200, "read file fail")
		return
	}

	{
		_, err := c.Writer.Write(data)
		if err != nil {
			return
		}
	}
}

// /unsafe/read_file2?fpath=../../../../../../etc/hosts
// /unsafe/read_file1?fpath=../../../../../../etc/hosts
