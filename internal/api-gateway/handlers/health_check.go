package handlers

import "github.com/gin-gonic/gin"

type IHealthCheck interface {
	HeathCheck(ctx *gin.Context)
}

type HeathCheckHandler struct{}

func NewHealthCheckHandler() IHealthCheck {
	return &HeathCheckHandler{}
}

func (h *HeathCheckHandler) HeathCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
