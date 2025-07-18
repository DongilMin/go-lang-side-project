package services

import (
	"errors"
	"shopping-mall-backend/models"
	"shopping-mall-backend/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	GetProductsByCategory(categoryID uint) ([]models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(id uint, product *models.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo  repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
}

func NewProductService(productRepo repositories.ProductRepository, categoryRepo repositories.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (s *productService) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	// Check if category exists
	_, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	
	return s.productRepo.GetByCategory(categoryID)
}

func (s *productService) CreateProduct(product *models.Product) error {
	// Validate category exists
	_, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}
	
	// Validate product data
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price < 0 {
		return errors.New("product price cannot be negative")
	}
	
	return s.productRepo.Create(product)
}

func (s *productService) UpdateProduct(id uint, product *models.Product) error {
	// Check if product exists
	existingProduct, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	
	// Validate category exists if changing category
	if product.CategoryID != existingProduct.CategoryID {
		_, err := s.categoryRepo.GetByID(product.CategoryID)
		if err != nil {
			return errors.New("category not found")
		}
	}
	
	// Update the existing product
	product.ID = id
	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	// Check if product exists
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New("product not found")
	}
	
	return s.productRepo.Delete(id)
}