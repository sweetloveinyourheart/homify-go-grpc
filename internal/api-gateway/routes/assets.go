package routes

import (
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/api-gateway/handlers"
	"homify-go-grpc/internal/api-gateway/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupAssetsHandler(
	router *gin.RouterGroup,
	client proto.PropertyClient,
	jwtAuthGuard *middlewares.JwtAuthGuard,
	validator *validator.Validate,
) {
	assetsHandler := handlers.NewAssetsHandler(client, validator)
	authGuard := jwtAuthGuard.AuthGuard

	// routes
	router.GET("/assets", assetsHandler.GetAssets)
	router.POST("/assets", authGuard, middlewares.RoleGuard("admin"), assetsHandler.AddNewAsset)
	router.PUT("/assets/modify", authGuard, middlewares.RoleGuard("admin"), assetsHandler.ModifyExistingAsset)
	router.PUT("/assets/disable", authGuard, middlewares.RoleGuard("admin"), assetsHandler.DisableAnAsset)
}
