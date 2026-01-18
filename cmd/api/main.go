package main

import (
	"log"
	"medical_master/internal/api"
	"medical_master/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB接続設定
	dsn := "host=localhost user=postgres password=postgres dbname=medical_master port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	r := gin.Default()

	// Swagger UI
	r.StaticFile("/doc/openapi.yaml", "./doc/openapi.yaml")
	config := &ginSwagger.Config{
		URL: "/doc/openapi.yaml",
	}
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// ルートアクセスを Swagger UI にリダイレクト
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// 自動生成されたハンドラーの登録
	serverImpl := &handler.ServerImpl{
		DB: db,
	}
	api.RegisterHandlers(r, serverImpl)

	log.Println("Server starting on http://localhost:8080")
	log.Println("Swagger UI available on http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
