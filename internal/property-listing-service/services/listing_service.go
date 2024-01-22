package services

import (
	"homify-go-grpc/internal/property-listing-service/models"
	"homify-go-grpc/internal/property-listing-service/repositories"

	"gorm.io/gorm"
)

type IPropertyListingService interface {
	AddNewProperty(newProperty models.Property) (bool, error)
}

type PropertyListingService struct {
	repo repositories.IPropertyRepository
}

func NewPropertyListingService(db *gorm.DB) IPropertyListingService {
	return &PropertyListingService{}
}

func (s *PropertyListingService) AddNewProperty(newProperty models.Property) (bool, error) {
	return true, nil
}
