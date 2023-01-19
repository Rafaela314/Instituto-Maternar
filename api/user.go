package api

import (
	"net/http"

	"github.com/Rafaela314/instituto-maternar/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {

	var req models.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: Add to DB
}
