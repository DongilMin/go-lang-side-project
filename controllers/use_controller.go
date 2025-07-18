package controllers

import (
    "net/http"
    "strconv"
    "shopping-mall-backend/database"
    "shopping-mall-backend/models"
    "github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    
    if err := database.DB.Preload("Orders").First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Create(&user)
    c.JSON(http.StatusCreated, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.DB.Save(&user)
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    database.DB.Delete(&models.User{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}