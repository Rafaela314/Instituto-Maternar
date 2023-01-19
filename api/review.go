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

type getReviewRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getReview(ctx *gin.Context) {

	var req getReviewRequest
	var err error

	err = ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: get from DB

	testReview, err := models.Review{}, nil
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, testReview)
}

type listReviewRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required, min=5, max=10"`
}

func (server *Server) listReview(ctx *gin.Context) {

	var req listReviewRequest
	var err error

	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: get from DB

	testReview, err := []models.Review{}, nil
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, testReview)
}
