package pkg

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"vuln-go-app/pkg/filter"
	"vuln-go-app/pkg/handler"
	"vuln-go-app/pkg/handler/research"
	"vuln-go-app/pkg/handler/safe"
	"vuln-go-app/pkg/handler/unsafe"
)

func InitRouter(e *gin.Engine) {
	e.Use(filter.Auth) // attention: middleware and "handler func" order

	e.GET("/ping", handler.Ping)
	e.POST("/user/login", handler.Login)

	// vulnerable
	e.GET("/unsafe/dig", unsafe.DigHost)
	e.GET("/unsafe/ssrf", unsafe.SSRF)
	e.GET("/unsafe/read_file1", unsafe.BadFileRead1)
	e.GET("/unsafe/read_file2", unsafe.BadFileRead2)
	e.POST("/unsafe/decompress_tar", unsafe.BadTarDecompress)
	e.GET("/unsafe/ssti1", unsafe.BadTemplate1)
	e.POST("/unsafe/ssti2", unsafe.BadTemplate2)

	// safe
	e.GET("/safe/fileread", safe.FileRead)
	e.GET("/safe/dig", safe.DigHost)
	e.POST("/safe/upload", safe.GoodUploadFile)

	// research
	e.GET("/research/realip", research.RealIP)
	e.GET("/research/panic", research.Panic)
	e.GET("/research/fatal_error", research.DeepRecursive)
	e.GET("/research/goodman", research.ConcurrentSecurity)
	e.GET("/research/mistake/:dir/*filename", research.MistakeCleanPath)
	e.POST("/research/http/read_body", research.ReadBody)
	e.GET("/research/http/read_body_flag", research.PrintFlag)

	//e.Static("/files/", "/etc/")
	e.Use(static.Serve("/files/", static.LocalFile("/etc", false)))
}
