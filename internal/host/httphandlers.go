package host

import (
	"net/http"

	"github.com/Rafaela314/instituto-maternar/models"
	"github.com/gin-gonic/gin"
)

// getReviewa responds with the list of all reviewa as JSON.
func (h *Host) GetReviews(c *gin.Context) {

	// albums slice to seed record album data.
	var albums = []models.Review{
		{ID: 1, Title: "Blue Train"},
		{ID: 2, Title: "Jeru"},
		{ID: 3, Title: "Sarah Vaughan and Clifford Brown"},
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// createReview responds with the list of all reviews as JSON.
func (h *Host) CreateReview(c *gin.Context) {

	var newReview models.Review

	err := c.BindJSON(&newReview)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, newReview)
}

// getReview responds with a review related to ID param as JSON.
func (h *Host) GetReviewByID(c *gin.Context) {

	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, id)
}
