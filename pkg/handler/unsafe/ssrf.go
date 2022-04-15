package unsafe

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}

func SSRF(c *gin.Context) {
	url := c.Query("url")
	c.Writer.Write(httpGet(url))
}

// /unsafe/ssrf?url=http://www.baidu.com
