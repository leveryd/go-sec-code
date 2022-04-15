package core

import (
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"vuln-go-app/pkg"
	"vuln-go-app/pkg/conf"
)

type Server struct {
	e *gin.Engine
}

func NewServer() *Server {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	s := &Server{}
	s.e = gin.New()
	//s.e.RemoveExtraSlash = true	// 影响路由匹配

	return s
}

func (s *Server) InitConfig() {
	//viper
}

func (s *Server) InitRouter() {
	pkg.InitRouter(s.e)
}

func (s *Server) Start() {

	addr := net.JoinHostPort(conf.ServerHost, strconv.Itoa(conf.ServerPort))
	s.e.Run(addr)

	//http.ListenAndServe(addr, nil)
}
