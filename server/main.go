package main

import (
	"fmt"
	"log"
	"missing-persons-backend/db"
	"missing-persons-backend/internal/handler"
	"missing-persons-backend/internal/repository"
	"missing-persons-backend/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	uploadDir := "./uploads/images"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()
	fmt.Println("Go server ready")

	baseURL := "https://missing-persons-databases-1.onrender.com"
	imageService := service.NewImageService(uploadDir,baseURL)

	personRepo := &repository.PersonRepository{
		DB: database,
	}

	personService := &service.PersonService{
		Repo: personRepo,
	}

	personHandler := &handler.PersonHandler{
		Service:      personService,
		ImageService: imageService,
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("api/persons", personHandler.GetMissingPersons)
	router.POST("api/persons", personHandler.CreateMissingPersons)
	router.Static("/images", uploadDir)

	fmt.Printf("Server starting on port :%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}