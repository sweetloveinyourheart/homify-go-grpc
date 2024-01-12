package api_gateway

import (
	"homify-go-grpc/internal/api-gateway/configs"
	"homify-go-grpc/internal/api-gateway/routes"

	"github.com/gin-gonic/gin"
)

func RunHTTPServer() {
	configurations := configs.GetConfig()

	router := gin.Default()
	routes.SetupHealthCheckRoute(router)

	if err := router.Run(configurations.Port); err != nil {
		panic(err)
	}
}
