package repositories

import (
	"homify-go-grpc/internal/property-service/models"

	"gorm.io/gorm"
)

// DestinationRepositoryInterface defines the interface for managing Destination entities.
type IDestinationRepository interface {
	CreateDestination(destination *models.Destination) error
	GetDestinationByID(id uint) (*models.Destination, error)
	GetAllCategories() ([]models.Destination, error)
	UpdateDestination(destination *models.Destination) error
	DeleteDestination(id uint) error
}

// DestinationRepository represents the repository for Destination.
type DestinationRepository struct {
	db *gorm.DB
}

// NewDestinationRepository creates a new DestinationRepository.
func NewDestinationRepository(db *gorm.DB) IDestinationRepository {
	return &DestinationRepository{db}
}

// CreateDestination creates a new destination.
func (r *DestinationRepository) CreateDestination(destination *models.Destination) error {
	return r.db.Create(destination).Error
}

// GetDestinationByID retrieves a destination by its ID.
func (r *DestinationRepository) GetDestinationByID(id uint) (*models.Destination, error) {
	var destination models.Destination
	err := r.db.First(&destination, id).Error
	return &destination, err
}

// GetAllCategories retrieves all amenities.
func (r *DestinationRepository) GetAllCategories() ([]models.Destination, error) {
	var categories []models.Destination
	err := r.db.Find(&categories).Error
	return categories, err
}

// UpdateDestination updates an existing destination.
func (r *DestinationRepository) UpdateDestination(destination *models.Destination) error {
	return r.db.Save(destination).Error
}

// DeleteDestination deletes a destination by its ID.
func (r *DestinationRepository) DeleteDestination(id uint) error {
	return r.db.Delete(&models.Destination{}, id).Error
}
