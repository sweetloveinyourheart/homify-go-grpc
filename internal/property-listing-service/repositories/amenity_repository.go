package repositories

import (
	"homify-go-grpc/internal/property-listing-service/models"

	"gorm.io/gorm"
)

// AmenityRepositoryInterface defines the interface for managing Amenity entities.
type IAmenityRepository interface {
	CreateAmenity(category *models.Amenity) error
	GetAmenityByID(id uint) (*models.Amenity, error)
	GetAllCategories() ([]models.Amenity, error)
	UpdateAmenity(category *models.Amenity) error
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

// CreateAmenity creates a new category.
func (r *AmenityRepository) CreateAmenity(category *models.Amenity) error {
	return r.db.Create(category).Error
}

// GetAmenityByID retrieves a category by its ID.
func (r *AmenityRepository) GetAmenityByID(id uint) (*models.Amenity, error) {
	var category models.Amenity
	err := r.db.First(&category, id).Error
	return &category, err
}

// GetAllCategories retrieves all categories.
func (r *AmenityRepository) GetAllCategories() ([]models.Amenity, error) {
	var categories []models.Amenity
	err := r.db.Find(&categories).Error
	return categories, err
}

// UpdateAmenity updates an existing category.
func (r *AmenityRepository) UpdateAmenity(category *models.Amenity) error {
	return r.db.Save(category).Error
}

// DeleteAmenity deletes a category by its ID.
func (r *AmenityRepository) DeleteAmenity(id uint) error {
	return r.db.Delete(&models.Amenity{}, id).Error
}
