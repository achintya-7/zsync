package server

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Define the route
	zsyncRouter := router.Group("/zsync")
	zsyncRouter.GET("/ping", )
	zsyncRouter.GET("/commands", )
	zsyncRouter.POST("/register", )

	server.router = router
}



