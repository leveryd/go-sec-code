package research

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var goodUser = false

func ConcurrentSecurity(c *gin.Context) {

	goodUser = false

	if c.Query("tell") != "" {
		// 模拟耗时io操作，可能导致协程调度
		//time.Sleep(time.Millisecond * 50)
		http.Get("http://www.baidu.com") // tell baidu i am good user
	}

	log.Println("goodUser:", goodUser)
	if goodUser == true {
		c.JSON(200, gin.H{
			"message": "good user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "bad user",
	})
}

func init() {
	go func() {
		log.Print("start check good user")
		for {
			//"good user?"
			goodUser = true

			// 模拟耗时io操作，可能导致协程调度
			// time.Sleep(time.Millisecond * 50)
			http.Get("http://www.baidu.com") // tell baidu i am good user

			//emm, bad user
			goodUser = false
		}
	}()
}

// 访问/research/goodman, 会一直提示bad user
// 访问/research/goodman?tell=1, 有可能会提示good user
