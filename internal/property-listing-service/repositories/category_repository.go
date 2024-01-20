package repositories

import (
	"homify-go-grpc/internal/property-listing-service/models"

	"gorm.io/gorm"
)

// CategoryRepositoryInterface defines the interface for managing Category entities.
type ICategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategoryByID(id uint) (*models.Category, error)
	GetAllCategories() ([]models.Category, error)
	UpdateCategory(category *models.Category) error
	DisableCategory(id uint) error
	DeleteCategory(id uint) error
}

// CategoryRepository represents the repository for Category.
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new CategoryRepository.
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{db}
}

// CreateCategory creates a new category.
func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

// GetCategoryByID retrieves a category by its ID.
func (r *CategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

// GetAllCategories retrieves all categories.
func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

// UpdateCategory updates an existing category.
func (r *CategoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

// DeleteCategory disables a category by its ID.
func (r *CategoryRepository) DisableCategory(id uint) error {
	var category models.Category

	err := r.db.First(&category, id).Error
	if err != nil {
		return err
	}

	category.IsAvailable = false

	return r.db.Save(category).Error
}

// DeleteCategory deletes a category by its ID.
func (r *CategoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
