package services

import (
	"errors"
	"shopping-mall-backend/models"
	"shopping-mall-backend/repositories"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(id uint, category *models.Category) error
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepo.GetAll()
}

func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	// Validate category data
	if category.Name == "" {
		return errors.New("category name is required")
	}
	
	return s.categoryRepo.Create(category)
}

func (s *categoryService) UpdateCategory(id uint, category *models.Category) error {
	// Check if category exists
	_, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	
	// Validate category data
	if category.Name == "" {
		return errors.New("category name is required")
	}
	
	// Update the existing category
	category.ID = id
	return s.categoryRepo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	// Check if category exists
	_, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	
	return s.categoryRepo.Delete(id)
}