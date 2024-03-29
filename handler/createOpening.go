package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrosandrini/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	openingCreateDTOResponse := schemas.OpeningCreateDTOResponse{
		ID:        opening.ID,
		CreatedAt: opening.CreatedAt,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
	}

	sendSuccess(ctx, "create-opening", openingCreateDTOResponse)
}
