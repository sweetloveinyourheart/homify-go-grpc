package handlers

import (
	proto "homify-go-grpc/api/property-listing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PropertyListingHandler struct {
	grpcClient proto.PropertyListingClient
	validator  *validator.Validate
}

func NewPropertyListingHandler(c proto.PropertyListingClient, validate *validator.Validate) *PropertyListingHandler {
	return &PropertyListingHandler{
		grpcClient: c,
		validator:  validate,
	}
}

func (h *PropertyListingHandler) AddNewProperty(ctx *gin.Context) {

}
