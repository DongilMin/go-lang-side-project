package repositories

import (
	"gorm.io/gorm"
	"shopping-mall-backend/models"
)

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Products").First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
