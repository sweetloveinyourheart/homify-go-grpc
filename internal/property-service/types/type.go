package types

import "homify-go-grpc/internal/property-service/models"

type PropertyAssets struct {
	Category models.Category
	Amenity  models.Amenity
}

type PropertyAssetIds struct {
	CategoryId uint
	AmenityId  uint
}
