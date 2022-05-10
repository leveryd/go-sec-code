package research

import (
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func GunzipHandler(c *gin.Context) {
	if strings.Contains(c.GetHeader("Content-Encoding"), "gzip") {
		reader, err := gzip.NewReader(c.Request.Body)
		if err != nil {
			return
		}
		defer reader.Close()

		{
			all, err := ioutil.ReadAll(reader)
			if err != nil {
				c.String(200, "read all error: %s", err.Error())
			}
			c.JSON(200, gin.H{
				"len": len(all),
			})
		}
	}
}

/*
dd if=/dev/zero of=/tmp/bigfile bs=512 count=200000
cat bigfile | gzip -c > /tmp/test.zip
*/

/*
import requests
with open("/tmp/test.zip", "rb") as f:
	requests.post("http://127.0.0.1:8089/research/http/unzip", headers={"Content-Encoding": "gzip"}, data=f.read())
*/

// https://stackoverflow.com/questions/56629115/how-to-protect-service-from-gzip-bomb
