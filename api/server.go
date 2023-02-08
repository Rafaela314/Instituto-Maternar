package api

import (
	db "github.com/Rafaela314/instituto-maternar/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.POST("/reviews", server.createReview)
	router.GET("/reviews/:id", server.getReview)
	router.GET("/reviews", server.listReview)
	router.POST("/users", server.createUser)

	server.router = router

	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
