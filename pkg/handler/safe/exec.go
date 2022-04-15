package safe

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
)

func DigHost(c *gin.Context) {
	host, exists := c.GetQuery("host")
	if exists == false {
		c.JSON(200, "no host arg")
		return
	}

	// safe
	command := "dig " + host
	commandResults, err := exec.Command(strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...).Output()
	//commandResults, err := exec.Command(command).Output()	//
	if err != nil {
		c.JSON(200, "fail to execute")
		return
	}
	c.Writer.Write(commandResults)
}
