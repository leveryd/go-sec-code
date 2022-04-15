package research

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func MistakeCleanPath(c *gin.Context) {

	prefix := os.TempDir()
	log.Println(prefix + "/example.txt")
	ioutil.WriteFile(prefix+"/example.txt", []byte("Hello World"), 0644)

	path := c.Param("dir") + c.Param("filename")
	absPath := filepath.Join(prefix, filepath.Clean(path))

	{
		file, err := ioutil.ReadFile(absPath)
		if err != nil {
			return
		}
		c.Writer.Write(file)
	}
}

// https://mp.weixin.qq.com/s/dqJ3F_fStlj78S0qhQ3Ggw

// /research/mistake/./../../../../../../../etc/hosts
// /research/mistake/./example
