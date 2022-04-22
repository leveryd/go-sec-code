package research

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var fileLock sync.Mutex

var tmpFilepath = fmt.Sprintf("%s/%s", os.TempDir(), "go-sec-code-tmp")

func ReadBody(c *gin.Context) {

	fileLock.Lock()
	defer fileLock.Unlock()

	file, err := os.OpenFile(tmpFilepath, os.O_RDWR|os.O_CREATE, 0640)
	if err != nil {
		return
	}
	defer os.Remove(tmpFilepath)

	hw := io.MultiWriter(file)
	io.Copy(hw, c.Request.Body)
}

func PrintFlag(c *gin.Context) {
	file, err := ioutil.ReadFile(tmpFilepath)
	if err != nil {
		return
	}
	if len(file) != 10 {
		return
	}

	time.Sleep(3 * time.Second) // check again after 3 second

	file, err = ioutil.ReadFile(tmpFilepath)
	if err != nil {
		return
	}
	if len(file) != 10 {
		return
	}

	c.JSON(200, "${ReadBody_flag}")
}

// https://www.leavesongs.com/PENETRATION/gitea-remote-command-execution.html
// doc/body读取.md

/*
curl 127.0.0.1:8089/research/http/read_body -d '0123456789' -H 'Content-Length:20'
curl 127.0.0.1:8089/research/http/read_body_flag
*/
