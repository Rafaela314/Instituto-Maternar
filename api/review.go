package api

import (
	"net/http"

	"github.com/Rafaela314/instituto-maternar/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) createReview(ctx *gin.Context) {

	var req models.Review
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: Add to DB
}
