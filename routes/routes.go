package routes

import (
	"shopping-mall-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, categoryController *controllers.CategoryController, productController *controllers.ProductController) {
	api := r.Group("/api/v1")

	// Category routes
	categories := api.Group("/categories")
	{
		categories.GET("", categoryController.GetCategories)
		categories.GET("/:id", categoryController.GetCategory)
		categories.POST("", categoryController.CreateCategory)
		categories.PUT("/:id", categoryController.UpdateCategory)
		categories.DELETE("/:id", categoryController.DeleteCategory)
		// 카테고리별 상품 조회를 여기에 배치
		categories.GET("/:id/products", productController.GetProductsByCategory)
	}

	// Product routes
	products := api.Group("/products")
	{
		products.GET("", productController.GetProducts)
		products.GET("/:id", productController.GetProduct)
		products.POST("", productController.CreateProduct)
		products.PUT("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)
	}
}