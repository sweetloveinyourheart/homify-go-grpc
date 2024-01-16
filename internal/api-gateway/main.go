package api_gateway

import (
	_ "homify-go-grpc/docs"
	"homify-go-grpc/internal/api-gateway/configs"
	"homify-go-grpc/internal/api-gateway/helpers"
	"homify-go-grpc/internal/api-gateway/routes"
	grpc_client "homify-go-grpc/internal/shared/grpc-client"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample Swagger API for a Go Gin application.
// @termsOfService https://example.com/terms/
// @contact name@example.com
// @license MIT

// @host localhost:8080
// @BasePath /api/v1
func RunHTTPServer() {
	configurations := configs.GetConfig()

	// Init auth client connection
	client, clientErr := grpc_client.NewGRPCAuthenticationClient(configurations.AuthenticationClientRemoteAddress)
	if clientErr != nil {
		panic(clientErr)
	}

	// Register the custom date validation function
	validator := validator.New()
	validator.RegisterValidation("customDate", helpers.CustomDate)

	router := gin.Default()

	// Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API group for version 1
	v1 := router.Group("/api/v1")
	{
		routes.SetupHealthCheckRoute(v1)
		routes.SetupAuthenticationHandler(v1, client, validator)
	}

	if err := router.Run(configurations.Port); err != nil {
		panic(err)
	}
}
