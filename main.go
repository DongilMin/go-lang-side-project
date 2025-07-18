package main

import (
	"shopping-mall-backend/controllers"
	"shopping-mall-backend/database"
	"shopping-mall-backend/models"
	"shopping-mall-backend/repositories"
	"shopping-mall-backend/routes"
	"shopping-mall-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 데이터베이스 연결 (하드코딩된 설정)
	database.Connect()

	// 테이블 마이그레이션
	database.DB.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.User{},
		&models.Order{},
		&models.OrderItem{},
		&models.CartItem{},
	)

	// Repository 초기화
	categoryRepo := repositories.NewCategoryRepository(database.DB)
	productRepo := repositories.NewProductRepository(database.DB)
	userRepo := repositories.NewUserRepository(database.DB)

	// Service 초기화
	categoryService := services.NewCategoryService(categoryRepo)
	productService := services.NewProductService(productRepo, categoryRepo)

	// Controller 초기화
	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)

	// Gin 라우터 설정
	r := gin.Default()

	// 라우트 설정 (의존성 주입)
	routes.SetupRoutes(r, categoryController, productController)

	// 서버 시작 (하드코딩된 포트)
	r.Run(":8080")
}