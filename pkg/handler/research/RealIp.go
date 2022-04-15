package research

import "github.com/gin-gonic/gin"

func RealIP(c *gin.Context) {
	//gin.SetMode(gin.ReleaseMode)
	c.Writer.Write([]byte(c.ClientIP()))
}

// https://github.com/gin-gonic/gin#dont-trust-all-proxies
