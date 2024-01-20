package repositories

import (
	"homify-go-grpc/internal/property-listing-service/models"

	"gorm.io/gorm"
)

// AmenityRepositoryInterface defines the interface for managing Amenity entities.
type IAmenityRepository interface {
	CreateAmenity(amenity *models.Amenity) error
	GetAmenityByID(id uint) (*models.Amenity, error)
	GetAllCategories() ([]models.Amenity, error)
	UpdateAmenity(amenity *models.Amenity) error
	DisableAmenity(id uint) error
	DeleteAmenity(id uint) error
}

// AmenityRepository represents the repository for Amenity.
type AmenityRepository struct {
	db *gorm.DB
}

// NewAmenityRepository creates a new AmenityRepository.
func NewAmenityRepository(db *gorm.DB) IAmenityRepository {
	return &AmenityRepository{db}
}

// CreateAmenity creates a new amenity.
func (r *AmenityRepository) CreateAmenity(amenity *models.Amenity) error {
	return r.db.Create(amenity).Error
}

// GetAmenityByID retrieves a amenity by its ID.
func (r *AmenityRepository) GetAmenityByID(id uint) (*models.Amenity, error) {
	var amenity models.Amenity
	err := r.db.First(&amenity, id).Error
	return &amenity, err
}

// GetAllCategories retrieves all amenities.
func (r *AmenityRepository) GetAllCategories() ([]models.Amenity, error) {
	var categories []models.Amenity
	err := r.db.Find(&categories).Error
	return categories, err
}

// UpdateAmenity updates an existing amenity.
func (r *AmenityRepository) UpdateAmenity(amenity *models.Amenity) error {
	return r.db.Save(amenity).Error
}

// DisableAmenity disables a amenity by its ID.
func (r *AmenityRepository) DisableAmenity(id uint) error {
	var amenity models.Category

	err := r.db.First(&amenity, id).Error
	if err != nil {
		return err
	}

	amenity.IsAvailable = false

	return r.db.Save(amenity).Error
}

// DeleteAmenity deletes a amenity by its ID.
func (r *AmenityRepository) DeleteAmenity(id uint) error {
	return r.db.Delete(&models.Amenity{}, id).Error
}
