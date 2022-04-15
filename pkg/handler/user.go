package handler

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	username string
	password string
}

func Login(c *gin.Context) {
	//body, err := c.Request.GetBody()
	//c.Request.Body
	//if err != nil {
	//	return
	//}

	//content := make([]byte, c.Request.ContentLength)
	//_, err := c.Request.Body.Read(content)
	//if err != nil {
	//	return
	//}

	//u := &User{}
	//json.Unmarshal(content, u)
	//
	//if u.username == "admin" && u.password == "admin" {
	//	c.Writer.Header().Set("username", "admin")
	//}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "admin" && password == "admin" {
		c.Writer.Header().Set("username", "admin")
	}
}
