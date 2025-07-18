package repositories

import (
	"shopping-mall-backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	GetByCategory(categoryID uint) ([]models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
	UpdateStock(id uint, quantity int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetByCategory(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) UpdateStock(id uint, quantity int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}
