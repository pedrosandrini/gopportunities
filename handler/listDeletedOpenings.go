package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrosandrini/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List Soft-Deleted openings
// @Description List all soft-deleted job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/deleted [get]
func ListDeletedOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error on list all deleted openings")
		return
	}

	sendSuccess(ctx, "list-deleted-openings", openings)
}
