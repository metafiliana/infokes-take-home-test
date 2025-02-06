package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	// Initialize database
	db = initDb()

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
	api := r.Group("/api")
	{
		// Example folder routes
		api.GET("/folders", handler.GetFolders)
		api.GET("/folders/:id", handler.GetSubFolders)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDb() *gorm.DB {
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
		"root", "", "localhost", "3306", "infokes-take-home-test")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
		panic(err)
	}

	return db
}
