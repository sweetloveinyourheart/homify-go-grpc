package handlers

import (
	"context"
	proto "homify-go-grpc/api/property"

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
	propertyCtx := context.Background()

	newProperty := &proto.NewProperty{
		Title: "Hi there !",
	}

	h.grpcClient.AddProperty(propertyCtx, newProperty)

	ctx.JSON(200, gin.H{
		"Success": true,
	})
}
