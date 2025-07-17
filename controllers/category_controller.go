package controllers

import (
	"net/http"
	"strconv"
	"shopping-mall-backend/database"
	"shopping-mall-backend/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category

	if err := database.DB.Preload("Products").First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&category)
	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category

	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Category{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}