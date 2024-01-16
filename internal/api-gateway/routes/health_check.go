package routes

import (
	"homify-go-grpc/internal/api-gateway/handlers"

	"github.com/gin-gonic/gin"
)

func SetupHealthCheckRoute(router *gin.RouterGroup) {
	heathCheckHandler := handlers.NewHealthCheckHandler()

	// routers
	router.GET("/health", heathCheckHandler.HeathCheck)
}
