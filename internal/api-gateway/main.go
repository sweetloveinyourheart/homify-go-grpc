package api_gateway

import (
	_ "homify-go-grpc/docs"
	"homify-go-grpc/internal/api-gateway/configs"
	"homify-go-grpc/internal/api-gateway/helpers"
	"homify-go-grpc/internal/api-gateway/middlewares"
	"homify-go-grpc/internal/api-gateway/routes"
	grpc_client "homify-go-grpc/internal/shared/grpc-client"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Homify API
// @version 1.0
// @description This is a sample Swagger API for a Go Gin application.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @produce json
// @SecurityDefinitions.apiKey Authorization
// @Security apiKey
// @in header
// @name Authorization
func RunHTTPServer() {
	configurations := configs.GetConfig()

	// Init auth client connection
	authClient, authClientErr := grpc_client.NewGRPCAuthenticationClient(configurations.AuthClientRemoteAddress)
	if authClientErr != nil {
		panic(authClientErr)
	}

	// Init property listing client connection
	propertyClient, propertyClientErr := grpc_client.NewGRPCPropertyClient(configurations.PropertyClientRemoteAddress)
	if propertyClientErr != nil {
		panic(propertyClientErr)
	}

	// Register the validation function
	validator := helpers.InitValidator()

	// Register the auth guard
	jwtAuthGuard := middlewares.NewJwtAuthGuard(authClient)

	router := gin.Default()

	// Swagger handler
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API group for version 1
	v1 := router.Group("/api/v1")
	{
		routes.SetupHealthCheckRoute(v1)
		routes.SetupAuthenticationHandler(v1, authClient, jwtAuthGuard, validator)
		routes.SetupAssetsHandler(v1, propertyClient, jwtAuthGuard, validator)
		routes.SetupPropertyRoute(v1, propertyClient, jwtAuthGuard, validator)
	}

	if err := router.Run(configurations.Port); err != nil {
		panic(err)
	}
}
