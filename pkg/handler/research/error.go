package research

import (
	"github.com/gin-gonic/gin"
)

func Panic(_ *gin.Context) {
	panic("test")
}

func DeepRecursive(c *gin.Context) {
	DeepRecursive(c)
}

// /research/fatal_error
// /research/panic

/*

net/http库有recover,所以项目中panic不会导致程序挂掉，但是"fatal error"会导致程序直接挂掉

json解析等功能会用到递归算法

*/

// [背事故？分享 6 种常见的 Go 致命错误场景](https://segmentfault.com/a/1190000041173313)
