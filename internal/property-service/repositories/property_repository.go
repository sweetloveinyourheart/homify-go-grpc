package repositories

import (
	"homify-go-grpc/internal/property-service/models"

	"gorm.io/gorm"
)

// PropertyRepositoryInterface defines the interface for managing Property entities.
type IPropertyRepository interface {
	CreateProperty(amenity *models.Property) error
	GetPropertyByID(id uint) (*models.Property, error)
	GetAllCategories() ([]models.Property, error)
	UpdateProperty(amenity *models.Property) error
	DeleteProperty(id uint) error
}

// PropertyRepository represents the repository for Property.
type PropertyRepository struct {
	db *gorm.DB
}

// NewPropertyRepository creates a new PropertyRepository.
func NewPropertyRepository(db *gorm.DB) IPropertyRepository {
	return &PropertyRepository{db}
}

// CreateProperty creates a new amenity.
func (r *PropertyRepository) CreateProperty(amenity *models.Property) error {
	return r.db.Create(amenity).Error
}

// GetPropertyByID retrieves a amenity by its ID.
func (r *PropertyRepository) GetPropertyByID(id uint) (*models.Property, error) {
	var amenity models.Property
	err := r.db.First(&amenity, id).Error
	return &amenity, err
}

// GetAllCategories retrieves all amenities.
func (r *PropertyRepository) GetAllCategories() ([]models.Property, error) {
	var categories []models.Property
	err := r.db.Find(&categories).Error
	return categories, err
}

// UpdateProperty updates an existing amenity.
func (r *PropertyRepository) UpdateProperty(amenity *models.Property) error {
	return r.db.Save(amenity).Error
}

// DeleteProperty deletes a amenity by its ID.
func (r *PropertyRepository) DeleteProperty(id uint) error {
	return r.db.Delete(&models.Property{}, id).Error
}
