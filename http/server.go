package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	client Handler
	engine *gin.Engine
}

func NewServer(client Handler) Server {
	return Server{
		client: client,
		engine: gin.Default(),
	}
}

func (s *Server) Run(port string) error {
	s.engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET, PUT, POST, DELETE, PATCH, OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Accept-Language"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	s.engine.GET("", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	client := s.engine.Group("/client")
	client.GET("", s.client.GetClients)
	client.GET("id", s.client.GetClientById)

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
