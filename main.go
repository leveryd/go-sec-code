package main

import (
	"vuln-go-app/pkg/core"
)

func main() {
	s := core.NewServer()
	s.InitConfig()
	s.InitRouter()
	s.Start()
}
