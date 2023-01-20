package server

import (
	"fmt"
	"go-crud/interface/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	address string
}

func StartApp(container container.Container) *Server {
	app := gin.Default()

	//setup handler
	handler := setupHandler(container)

	//setup router
	setUpRouter(app, handler)

	port := ":8888"

	return &Server{
		engine:  app,
		address: port,
	}
}

func (s *Server) Run() {
	fmt.Println(s.engine.Run(s.address))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.engine.ServeHTTP(w, r)
}
