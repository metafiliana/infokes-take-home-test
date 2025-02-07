package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/infokes-take-home-test/backend/configs"
	"github.com/infokes-take-home-test/backend/handlers"
	"github.com/infokes-take-home-test/backend/repositories"
	"github.com/infokes-take-home-test/backend/usecases"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func main() {
	// init env config variable
	cfg := configs.Get()

	// Initialize database
	db = initDb(cfg)

	// Create Gin router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize repositories
	folderRepo := repositories.NewFolderRepository(db)

	//Initialize usecase
	folderUsecase := usecases.NewFolderUsecase(folderRepo)

	// Initialize handler
	handler := handlers.NewFolderHandler(folderUsecase)

	// Group routes
	v1 := r.Group("/api/v1")
	{
		// Example folder routes
		v1.GET("/folders", handler.GetFolders)
		v1.GET("/folders/:id", handler.GetSubFolders)
	}

	// Start server
	fmt.Printf("Server starting on port %s\n", cfg.RestPort)
	if err := r.Run(":" + cfg.RestPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDb(cfg configs.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	// local mode development
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
		panic(err)
	}

	return db
}
