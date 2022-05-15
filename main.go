package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vuln-go-app/pkg/core"
)

func main() {
	s := core.NewServer()

	go func() {
		s.InitConfig()
		s.InitRouter()
		s.Start()
	}()

	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit

	log.Println("Shutdown Server ...")
	// 调用Server.Shutdown graceful结束
	if err := s.Server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
