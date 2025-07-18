package routes

import (
    "shopping-mall-backend/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api/v1")
    
    // Category routes
    categories := api.Group("/categories")
    {
        categories.GET("", controllers.GetCategories)
        categories.GET("/:id", controllers.GetCategory)
        categories.POST("", controllers.CreateCategory)
        categories.PUT("/:id", controllers.UpdateCategory)
        categories.DELETE("/:id", controllers.DeleteCategory)
    }
    
    // Product routes
    products := api.Group("/products")
    {
        products.GET("", controllers.GetProducts)
        products.GET("/:id", controllers.GetProduct)
        products.POST("", controllers.CreateProduct)
        products.PUT("/:id", controllers.UpdateProduct)
        products.DELETE("/:id", controllers.DeleteProduct)
    }
    
    // User routes
    users := api.Group("/users")
    {
        users.GET("", controllers.GetUsers)
        users.GET("/:id", controllers.GetUser)
        users.POST("", controllers.CreateUser)
        users.PUT("/:id", controllers.UpdateUser)
        users.DELETE("/:id", controllers.DeleteUser)
    }
}