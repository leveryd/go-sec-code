package safe

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"regexp"
)

func DigHost(c *gin.Context) {
	host, exists := c.GetQuery("host")
	if exists == false {
		c.JSON(200, "no host arg")
		return
	}

	// safe: whitelist
	{
		compile, err := regexp.Compile(`^([\w\d.\-_]+)$`)
		if err != nil {
			return
		}
		if compile.Match([]byte(host)) == false {
			c.JSON(200, "invalid host arg")
			return
		}
	}

	commandResults, err := exec.Command("dig", host).Output()

	if err != nil {
		c.JSON(200, "fail to execute")
		return
	}
	c.Writer.Write(commandResults)
}

// /safe/dig?host=www.baidu.com

// https://github.com/leveryd/go-sec-code/issues/2
// 即使没有使用bash -c，也需要注意命令可能有参数能够被用来实施攻击

// 或许需要补充其他类型的安全写法，除了黑名单
