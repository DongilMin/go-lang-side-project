package controllers

import (
    "net/http"
    "strconv"
    "shopping-mall-backend/database"
    "shopping-mall-backend/models"
    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    var products []models.Product
    database.DB.Preload("Category").Find(&products)
    c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var product models.Product
    
    if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Create(&product)
    c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var product models.Product
    
    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Save(&product)
    c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    database.DB.Delete(&models.Product{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}