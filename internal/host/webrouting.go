package host

import (
	"github.com/gin-gonic/gin"
)

// CreateWebRouter creates an HttpRouter and maps all the endpoints to handlers.
func (h *Host) CreateWebRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/reviews", h.GetReviews)
	router.GET("/reviews/:id", h.GetReviewByID)
	router.POST("/reviews", h.CreateReview)

	router.Run("localhost:8080")

	return router
}
