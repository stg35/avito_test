package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stg35/avito_test/internal/config"
	"github.com/stg35/avito_test/internal/handler"
)

type Server struct {
	router *gin.Engine
}

func NewServer(handler *handler.Handler) *Server {
	router := gin.New()
	router.Use(gin.Logger())
	handler.InitRoutes(router)

	return &Server{
		router,
	}
}

func (server *Server) Start(config *config.ServerConfig) error {
	fmt.Println(config.Host, config.Port)
	return server.router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
