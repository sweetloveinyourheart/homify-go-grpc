package services

import "gorm.io/gorm"

type IPropertyListingService interface{}

type PropertyListingService struct{}

func NewPropertyListingService(db *gorm.DB) IPropertyListingService {
	return &PropertyListingService{}
}
