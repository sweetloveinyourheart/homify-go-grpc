package server

import (
	"context"
	"fmt"
	proto "homify-go-grpc/api/property"
	"homify-go-grpc/internal/property-service/models"
	"homify-go-grpc/internal/property-service/producers"
	"homify-go-grpc/internal/property-service/services"
	"homify-go-grpc/internal/property-service/types"

	"gorm.io/gorm"
)

type GRPCPropertyServer struct {
	amenitySvc  services.IAmenityService
	categorySvc services.ICategoryService
	propertySvc services.IPropertyService
	proto.UnimplementedPropertyServer
}

func NewGRPCPropertyServer(db *gorm.DB, p producers.IPropertyProducer) *GRPCPropertyServer {
	amenitySvc := services.NewAmenityService(db)
	categorySvc := services.NewCategoryService(db)
	propertySvc := services.NewPropertyService(db, p)

	return &GRPCPropertyServer{
		amenitySvc:  amenitySvc,
		categorySvc: categorySvc,
		propertySvc: propertySvc,
	}
}

const (
	CATEGORY_ASSET_TYPE = "category"
	AMENITY_ASSET_TYPE  = "amenity"
)

// Property handlers

func (s *GRPCPropertyServer) AddProperty(ctx context.Context, req *proto.NewProperty) (*proto.ResultResponse, error) {
	assetIds := types.PropertyAssetIds{
		CategoryId: uint(req.CategoryId),
		AmenityId:  uint(req.AmenityId),
	}

	newProperty := models.Property{
		HostId:      uint(req.HostId),
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
	}

	newDestination := models.Destination{
		Country:   req.Destination.Country,
		City:      req.Destination.City,
		Latitude:  req.Destination.Latitude,
		Longitude: req.Destination.Longitude,
	}

	s.propertySvc.AddNewProperty(assetIds, newProperty, newDestination)
	return &proto.ResultResponse{Success: true}, nil
}

// Assets handlers

func (s *GRPCPropertyServer) GetAssets(ctx context.Context, req *proto.GetAssetsRequest) (*proto.GetAssetsResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		assets, err := s.categorySvc.GetCategories()

		if err != nil {
			return &proto.GetAssetsResponse{}, err
		}

		// Convert models.Category
		var convertedAssets []*proto.Assets
		for _, asset := range assets {
			convertedAsset := &proto.Assets{
				// Map fields from models.Category
				Id:      uint32(asset.ID),
				IconURL: asset.IconURL,
				Name:    asset.Name,
			}
			convertedAssets = append(convertedAssets, convertedAsset)
		}

		return &proto.GetAssetsResponse{
			Assets: convertedAssets,
		}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		assets, err := s.amenitySvc.GetAmenities()

		if err != nil {
			return &proto.GetAssetsResponse{}, err
		}

		// Convert models.Amenity
		var convertedAssets []*proto.Assets
		for _, asset := range assets {
			convertedAsset := &proto.Assets{
				// Map fields from models.Amenity
				Id:      uint32(asset.ID),
				IconURL: asset.IconURL,
				Name:    asset.Name,
			}
			convertedAssets = append(convertedAssets, convertedAsset)
		}

		return &proto.GetAssetsResponse{
			Assets: convertedAssets,
		}, nil
	}

	return &proto.GetAssetsResponse{}, nil
}

func (s *GRPCPropertyServer) AddAsset(ctx context.Context, req *proto.AddAssetRequest) (*proto.ResultResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		err := s.categorySvc.CreateCategory(&models.Category{
			Name:    req.Name,
			IconURL: req.IconURL,
		})

		if err != nil {
			return &proto.ResultResponse{Success: false}, err
		}

		return &proto.ResultResponse{Success: true}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		err := s.amenitySvc.CreateAmenity(&models.Amenity{
			Name:    req.Name,
			IconURL: req.IconURL,
		})

		if err != nil {
			return &proto.ResultResponse{Success: false}, err
		}

		return &proto.ResultResponse{Success: true}, nil
	}

	return &proto.ResultResponse{Success: false}, fmt.Errorf("invalid asset type")
}

func (s *GRPCPropertyServer) ModifyAsset(ctx context.Context, req *proto.ModifyAssetRequest) (*proto.ResultResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		err := s.categorySvc.UpdateCategory(
			uint(req.Id),
			&models.Category{
				Name:    req.Name,
				IconURL: req.IconURL,
			},
		)

		if err != nil {
			return &proto.ResultResponse{Success: false}, err
		}

		return &proto.ResultResponse{Success: true}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		err := s.amenitySvc.UpdateAmenity(
			uint(req.Id),
			&models.Amenity{
				Name:    req.Name,
				IconURL: req.IconURL,
			},
		)

		if err != nil {
			return &proto.ResultResponse{Success: false}, err
		}

		return &proto.ResultResponse{Success: true}, nil
	}

	return &proto.ResultResponse{Success: false}, fmt.Errorf("invalid asset type")
}

func (s *GRPCPropertyServer) DisableAsset(ctx context.Context, req *proto.DisableAssetRequest) (*proto.ResultResponse, error) {
	if req.AssetType == CATEGORY_ASSET_TYPE {
		err := s.categorySvc.DisableCategory(uint(req.Id))

		if err != nil {
			return &proto.ResultResponse{Success: false}, err
		}

		return &proto.ResultResponse{Success: true}, nil
	}

	if req.AssetType == AMENITY_ASSET_TYPE {
		err := s.amenitySvc.DisableAmenity(uint(req.Id))

		if err != nil {
			return &proto.ResultResponse{Success: false}, err
		}

		return &proto.ResultResponse{Success: true}, nil
	}

	return &proto.ResultResponse{Success: false}, fmt.Errorf("invalid asset type")
}
