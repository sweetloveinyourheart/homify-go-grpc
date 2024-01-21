package services

import (
	"homify-go-grpc/internal/property-listing-service/models"
	"homify-go-grpc/internal/property-listing-service/repositories"

	"gorm.io/gorm"
)

type IAmenityService interface {
	CreateAmenity(amenity *models.Amenity) error
	GetAmenities() ([]models.Amenity, error)
	GetAmenityByID(amenityID uint) (*models.Amenity, error)
	UpdateAmenity(amenityID uint, amenity *models.Amenity) error
	DisableAmenity(amenityID uint) error
	DeleteAmenity(amenityID uint) error
}

type AmenityService struct {
	repo repositories.IAmenityRepository
}

func NewAmenityService(db *gorm.DB) IAmenityService {
	return &AmenityService{
		repo: repositories.NewAmenityRepository(db),
	}
}

func (s *AmenityService) GetAmenities() ([]models.Amenity, error) {
	return s.repo.GetAllCategories()
}

func (s *AmenityService) CreateAmenity(amenity *models.Amenity) error {
	return s.repo.CreateAmenity(amenity)
}

func (s *AmenityService) GetAmenityByID(amenityID uint) (*models.Amenity, error) {
	return s.repo.GetAmenityByID(amenityID)
}

func (s *AmenityService) UpdateAmenity(amenityID uint, updateData *models.Amenity) error {
	// Check if the amenity exists
	amenity, err := s.GetAmenityByID(amenityID)
	if err != nil {
		return err // Amenity not found
	}

	if updateData.Name != "" {
		amenity.Name = updateData.Name
	}

	if updateData.IconURL != "" {
		amenity.IconURL = updateData.IconURL
	}

	amenity.ID = amenityID

	return s.repo.UpdateAmenity(amenity)
}

func (s *AmenityService) DisableAmenity(amenityID uint) error {
	// Check if the amenity exists
	amenity, err := s.GetAmenityByID(amenityID)
	if err != nil {
		return err // Amenity not found
	}

	amenity.IsAvailable = false

	return s.repo.UpdateAmenity(amenity)
}

func (s *AmenityService) DeleteAmenity(amenityID uint) error {
	// Check if the amenity exists
	_, err := s.GetAmenityByID(amenityID)
	if err != nil {
		return err // Amenity not found
	}

	return s.repo.DeleteAmenity(amenityID)
}
