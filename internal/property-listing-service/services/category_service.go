package services

import (
	"homify-go-grpc/internal/property-listing-service/models"
	"homify-go-grpc/internal/property-listing-service/repositories"

	"gorm.io/gorm"
)

type ICategoryService interface {
	CreateCategory(category *models.Category) error
	GetCategories() ([]models.Category, error)
	GetCategoryByID(categoryID uint) (*models.Category, error)
	UpdateCategory(categoryID uint, updateData *models.Category) error
	DisableCategory(categoryID uint) error
	DeleteCategory(categoryID uint) error
}

type CategoryService struct {
	repo repositories.ICategoryRepository
}

func NewCategoryService(db *gorm.DB) ICategoryService {
	return &CategoryService{
		repo: repositories.NewCategoryRepository(db),
	}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) GetCategoryByID(categoryID uint) (*models.Category, error) {
	return s.repo.GetCategoryByID(categoryID)
}

func (s *CategoryService) UpdateCategory(categoryID uint, updateData *models.Category) error {
	// Check if the category exists
	category, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err // Category not found
	}

	if updateData.Name != "" {
		category.Name = updateData.Name
	}

	if updateData.IconURL != "" {
		category.IconURL = updateData.IconURL
	}

	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DisableCategory(categoryID uint) error {
	// Check if the category exists
	category, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err // Category not found
	}

	category.IsAvailable = false

	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(categoryID uint) error {
	// Check if the category exists
	_, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err // Category not found
	}

	return s.repo.DeleteCategory(categoryID)
}
