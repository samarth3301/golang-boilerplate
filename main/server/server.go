package server

import (
	"context"
	"fmt"
	"golang-boilerplate/main/config"
	"golang-boilerplate/main/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	server *http.Server
}

func NewServer() *Server {
	router := gin.Default()
	routes.SetupRoutes(router)

	return &Server{
		router: router,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.AppConfig.Server.Host, config.AppConfig.Server.Port),
			Handler: router,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
