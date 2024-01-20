package services

import (
	"homify-go-grpc/internal/property-listing-service/models"
	"homify-go-grpc/internal/property-listing-service/repositories"

	"gorm.io/gorm"
)

type ICategoryService interface {
	CreateCategory(category *models.Category) error
	GetCategoryByID(categoryID uint) (*models.Category, error)
	UpdateCategory(category *models.Category) error
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

func (s *CategoryService) GetCategoryByID(categoryID uint) (*models.Category, error) {
	return s.repo.GetCategoryByID(categoryID)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	// Check if the category exists
	_, err := s.GetCategoryByID(category.ID)
	if err != nil {
		return err // Category not found
	}

	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DisableCategory(categoryID uint) error {
	// Check if the category exists
	_, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err // Category not found
	}

	return s.repo.DisableCategory(categoryID)
}

func (s *CategoryService) DeleteCategory(categoryID uint) error {
	// Check if the category exists
	_, err := s.GetCategoryByID(categoryID)
	if err != nil {
		return err // Category not found
	}

	return s.repo.DeleteCategory(categoryID)
}
