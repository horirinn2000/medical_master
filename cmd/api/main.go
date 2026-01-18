package main

import (
	"context"
	"fmt"
	"log"
	"medical_master/internal/api"
	"medical_master/internal/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// DB接続設定
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "medical_master"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
		getEnv("DB_TIMEZONE", "Asia/Tokyo"),
	)
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

	// api.RegisterHandlers は api.ServerInterface を受け取る
	// ServerImpl が必要なメソッドをすべて実装しているか確認
	serverImpl := handler.NewServerImpl(db)
	api.RegisterHandlers(r, serverImpl)

	port := getEnv("PORT", "8080")
	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("Swagger UI available on http://localhost:%s/swagger/index.html", port)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// サーバー起動（goroutineで）
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 終了シグナル待機（省略可能、必要に応じて実装）
	// シグナル待ち
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Graceful Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get db from gorm: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}

	log.Println("Server exiting")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
