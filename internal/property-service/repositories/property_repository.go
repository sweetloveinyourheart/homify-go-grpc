package repositories

import (
	"homify-go-grpc/internal/property-service/models"

	"gorm.io/gorm"
)

// PropertyRepositoryInterface defines the interface for managing Property entities.
type IPropertyRepository interface {
	CreateProperty(property *models.Property) error
	GetPropertyByID(id uint) (*models.Property, error)
	GetAllCategories() ([]models.Property, error)
	Association(property *models.Property, column string) *gorm.Association
	UpdateProperty(property *models.Property) error
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

// CreateProperty creates a new property.
func (r *PropertyRepository) CreateProperty(property *models.Property) error {
	return r.db.Create(property).Error
}

// GetPropertyByID retrieves a property by its ID.
func (r *PropertyRepository) GetPropertyByID(id uint) (*models.Property, error) {
	var property models.Property
	err := r.db.First(&property, id).Error
	return &property, err
}

// GetAllCategories retrieves all amenities.
func (r *PropertyRepository) GetAllCategories() ([]models.Property, error) {
	var categories []models.Property
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *PropertyRepository) Association(property *models.Property, column string) *gorm.Association {
	return r.db.Model(property).Association(column)
}

// UpdateProperty updates an existing property.
func (r *PropertyRepository) UpdateProperty(property *models.Property) error {
	return r.db.Save(property).Error
}

// DeleteProperty deletes a property by its ID.
func (r *PropertyRepository) DeleteProperty(id uint) error {
	return r.db.Delete(&models.Property{}, id).Error
}
