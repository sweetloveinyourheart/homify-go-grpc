package handlers

import (
	"github.com/gin-gonic/gin"
)

type IHealthCheck interface {
	HeathCheck(ctx *gin.Context)
}

type HeathCheckHandler struct{}

type HeathCheckResponse struct {
	Message string `json:"message"`
}

func NewHealthCheckHandler() IHealthCheck {
	return &HeathCheckHandler{}
}

// HeathCheck godoc
// @Tags Health check
// @Summary Health check endpoint
// @Description Returns a simple health check response indicating the status of the service.
// @Produce json
// @Success 200 {object} HeathCheckResponse

// @Router /health [get]
func (h *HeathCheckHandler) HeathCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
