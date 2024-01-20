package server

import (
	"context"
	"fmt"
	proto "homify-go-grpc/api/property-listing"
	"homify-go-grpc/internal/property-listing-service/models"
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

const (
	CATEGORY_ASSET_TYPE = "category"
	AMENITY_ASSET_TYPE  = "amenity"
)

func (s *GRPCPropertyListingServer) AddAsset(ctx context.Context, req *proto.AddAssetRequest) (*proto.AddAssetResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		err := s.categorySvc.CreateCategory(&models.Category{
			Name:    req.Name,
			IconURL: req.IconURL,
		})

		if err != nil {
			return &proto.AddAssetResponse{Success: false}, err
		}

		return &proto.AddAssetResponse{Success: true}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		err := s.amenitySvc.CreateAmenity(&models.Amenity{
			Name:    req.Name,
			IconURL: req.IconURL,
		})

		if err != nil {
			return &proto.AddAssetResponse{Success: false}, err
		}

		return &proto.AddAssetResponse{Success: true}, nil
	}

	return &proto.AddAssetResponse{Success: false}, fmt.Errorf("invalid asset type")
}

func (s *GRPCPropertyListingServer) ModifyAsset(ctx context.Context, req *proto.ModifyAssetRequest) (*proto.ModifyAssetResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		err := s.categorySvc.UpdateCategory(&models.Category{
			Name:    req.Name,
			IconURL: req.IconURL,
		})

		if err != nil {
			return &proto.ModifyAssetResponse{Success: false}, err
		}

		return &proto.ModifyAssetResponse{Success: true}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		err := s.amenitySvc.UpdateAmenity(&models.Amenity{
			Name:    req.Name,
			IconURL: req.IconURL,
		})

		if err != nil {
			return &proto.ModifyAssetResponse{Success: false}, err
		}

		return &proto.ModifyAssetResponse{Success: true}, nil
	}

	return &proto.ModifyAssetResponse{Success: false}, fmt.Errorf("invalid asset type")
}

func (s *GRPCPropertyListingServer) DisableAsset(ctx context.Context, req *proto.DisableAssetRequest) (*proto.DisableAssetResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		err := s.categorySvc.DisableCategory(uint(req.Id))

		if err != nil {
			return &proto.DisableAssetResponse{Success: false}, err
		}

		return &proto.DisableAssetResponse{Success: true}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		err := s.amenitySvc.DisableAmenity(uint(req.Id))

		if err != nil {
			return &proto.DisableAssetResponse{Success: false}, err
		}

		return &proto.DisableAssetResponse{Success: true}, nil
	}

	return &proto.DisableAssetResponse{Success: false}, fmt.Errorf("invalid asset type")
}
