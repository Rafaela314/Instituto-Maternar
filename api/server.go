package api

import "github.com/gin-gonic/gin"

// Server serves HTTP requests
type Server struct {
	// TODO: add store *db.Dtore for database interaction
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer() *Server {
	server := &Server{}

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
