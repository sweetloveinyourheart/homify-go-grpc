package services

import (
	"homify-go-grpc/internal/property-listing-service/models"
	"homify-go-grpc/internal/property-listing-service/repositories"

	"gorm.io/gorm"
)

type IAmenityService interface {
	CreateAmenity(category *models.Amenity) error
	GetAmenityByID(categoryID uint) (*models.Amenity, error)
	UpdateAmenity(category *models.Amenity) error
	DisableAmenity(categoryID uint) error
	DeleteAmenity(categoryID uint) error
}

type AmenityService struct {
	repo repositories.IAmenityRepository
}

func NewAmenityService(db *gorm.DB) IAmenityService {
	return &AmenityService{
		repo: repositories.NewAmenityRepository(db),
	}
}

func (s *AmenityService) CreateAmenity(category *models.Amenity) error {
	return s.repo.CreateAmenity(category)
}

func (s *AmenityService) GetAmenityByID(categoryID uint) (*models.Amenity, error) {
	return s.repo.GetAmenityByID(categoryID)
}

func (s *AmenityService) UpdateAmenity(category *models.Amenity) error {
	// Check if the category exists
	_, err := s.GetAmenityByID(category.ID)
	if err != nil {
		return err // Amenity not found
	}

	return s.repo.UpdateAmenity(category)
}

func (s *AmenityService) DisableAmenity(categoryID uint) error {
	// Check if the category exists
	_, err := s.GetAmenityByID(categoryID)
	if err != nil {
		return err // Amenity not found
	}

	return s.repo.DisableAmenity(categoryID)
}

func (s *AmenityService) DeleteAmenity(categoryID uint) error {
	// Check if the category exists
	_, err := s.GetAmenityByID(categoryID)
	if err != nil {
		return err // Amenity not found
	}

	return s.repo.DeleteAmenity(categoryID)
}
