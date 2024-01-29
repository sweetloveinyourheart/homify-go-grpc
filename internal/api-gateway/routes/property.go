package routes

import (
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/api-gateway/handlers"
	"homify-go-grpc/internal/api-gateway/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupPropertyRoute(
	router *gin.RouterGroup,
	client proto.PropertyClient,
	jwtAuthGuard *middlewares.JwtAuthGuard,
	validator *validator.Validate,
) {
	propertyHandler := handlers.NewPropertyHandler(client, validator)
	authGuard := jwtAuthGuard.AuthGuard

	// routers
	router.POST("/property", authGuard, propertyHandler.AddNewProperty)
	router.PUT("/property/sync", propertyHandler.SyncProperties)
}
