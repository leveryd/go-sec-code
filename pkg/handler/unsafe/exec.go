package unsafe

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os/exec"
)

func DigHost(c *gin.Context) {
	host, exists := c.GetQuery("host")
	if exists == false {
		c.JSON(200, "no host arg")
		return
	}

	//cmd := fmt.Sprintf("/bin/bash -c 'dig %s'", host)
	cmd := fmt.Sprintf("dig %s", host)
	log.Print(cmd)

	commandResults, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		c.JSON(200, "fail to execute")
		return
	}
	c.Writer.Write(commandResults)
}

// /unsafe/dig?host=www.baidu.com%26%26id
