package main

import (
    "shopping-mall-backend/database"
    "shopping-mall-backend/models"
    "shopping-mall-backend/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // 데이터베이스 연결
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
    
    // Gin 라우터 설정
    r := gin.Default()
    
    // 라우트 설정
    routes.SetupRoutes(r)
    
    // 서버 시작
    r.Run(":8080")
}