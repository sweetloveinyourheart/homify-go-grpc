package server

import (
	proto "homify-go-grpc/api/property-listing"
	"homify-go-grpc/internal/property-listing-service/services"

	"gorm.io/gorm"
)

type GRPCPropertyListingServer struct {
	amenitySvc  services.IAmenityService
	categorySvc services.ICategoryService
	listingSvc  services.IPropertyListingService
	proto.UnimplementedPropertyListingServer
}

func NewGRPCPropertyListingServer(db *gorm.DB) *GRPCPropertyListingServer {
	amenitySvc := services.NewAmenityService(db)
	categorySvc := services.NewCategoryService(db)
	listingSvc := services.NewPropertyListingService(db)

	return &GRPCPropertyListingServer{
		amenitySvc:  amenitySvc,
		categorySvc: categorySvc,
		listingSvc:  listingSvc,
	}
}
