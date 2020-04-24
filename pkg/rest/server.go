package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhiqiangxu/ncrypt/pkg/instance"
	"github.com/zhiqiangxu/ncrypt/pkg/rest/api"
	"go.uber.org/zap"
)

// Server model
type Server struct {
	app *gin.Engine
}

// NewServer creates a Server
func NewServer() *Server {

	s := &Server{}
	app := api.New()

	s.app = app

	return s
}

// Start the server
func (s *Server) Start(addr string) {

	httpServer := &http.Server{Addr: addr, Handler: s.app, ReadTimeout: time.Second * 3}
	err := httpServer.ListenAndServe()
	if err != nil {
		instance.Logger().Error("ListenAndServe", zap.Error(err))
		return
	}
}
