package routes

import (
	proto "homify-go-grpc/api/authentication"
	"homify-go-grpc/internal/api-gateway/handlers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupAuthenticationHandler(router *gin.RouterGroup, client proto.AuthenticationClient, validator *validator.Validate) {
	authenticationHandler := handlers.NewAuthHandler(client, validator)

	// routes
	router.POST("/sign-up", authenticationHandler.SignUp)
	router.POST("/sign-in", authenticationHandler.SignIn)
}
