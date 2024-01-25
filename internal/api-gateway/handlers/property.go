package handlers

import (
	"context"
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/api-gateway/dtos"
	"homify-go-grpc/internal/api-gateway/helpers"
	"homify-go-grpc/internal/api-gateway/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PropertyHandler struct {
	grpcClient proto.PropertyClient
	validator  *validator.Validate
}

func NewPropertyHandler(c proto.PropertyClient, validate *validator.Validate) *PropertyHandler {
	return &PropertyHandler{
		grpcClient: c,
		validator:  validate,
	}
}

// @Summary Add a new property
// @Tags Property
// @Description Add a new property to the property listings.
// @ID add-new-property
// @Accept  json
// @Produce  json
// @Security Authorization
// @Param input body dtos.NewPropertyDTO true "New Property object to be added"
// @Success 200 {object} interface{} "Successfully added the new property"
// @Router /property [post]
func (h *PropertyHandler) AddNewProperty(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(403, gin.H{"error": "Forbidden resource"})
		return
	}

	authenticatedUser := user.(middlewares.AuthenticatedUser)

	newPropertyData := dtos.NewPropertyDTO{}
	if bindError := ctx.ShouldBindJSON(&newPropertyData); bindError != nil {
		ctx.JSON(400, gin.H{"error": bindError.Error()})
		return
	}

	validatorError := h.validator.Struct(newPropertyData)
	if validatorError != nil {
		helpers.HandleValidationErrors(ctx, validatorError)
		return
	}

	newProperty := &proto.NewProperty{
		HostId:      uint32(authenticatedUser.Id),
		Title:       newPropertyData.Title,
		Description: newPropertyData.Description,
		Price:       newPropertyData.Price,
		AmenityId:   uint32(newPropertyData.AmenityId),
		CategoryId:  uint32(newPropertyData.CategoryId),
		Destination: &proto.NewDestination{
			Country:   newPropertyData.Destination.Country,
			City:      newPropertyData.Destination.City,
			Latitude:  newPropertyData.Destination.Latitude,
			Longitude: newPropertyData.Destination.Longitude,
		},
	}

	propertyCtx := context.Background()

	h.grpcClient.AddProperty(propertyCtx, newProperty)

	ctx.JSON(200, gin.H{
		"Success": true,
	})
}
