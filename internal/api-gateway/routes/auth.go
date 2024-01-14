package routes

import (
	"homify-go-grpc/internal/api-gateway/handlers"

	"github.com/gin-gonic/gin"
)

var authenticationHandler = handlers.NewAuthHandler()

func SetupAuthenticationHandler(router *gin.RouterGroup) {
	router.POST("/sign-up", authenticationHandler.SignUp)
}
