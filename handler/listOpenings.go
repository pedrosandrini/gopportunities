package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrosandrini/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error on list openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
