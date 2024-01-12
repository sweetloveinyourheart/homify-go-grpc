package routes

import (
	"homify-go-grpc/internal/api-gateway/handlers"

	"github.com/gin-gonic/gin"
)

var heathCheckHandler = handlers.NewHealthCheckHandler()

func SetupHealthCheckRoute(router *gin.Engine) {
	router.GET("/health", heathCheckHandler.HeathCheck)
}
