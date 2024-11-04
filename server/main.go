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

	

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()
	fmt.Println("Go server ready")

	// baseURL := "https://missing-persons-databases-1.onrender.com"
	imageService, err := service.NewImageService(
		os.Getenv("S3_BUCKET_NAME"),
		fmt.Sprintf("https://%s.s3.%s.amazonaws.com", 
			os.Getenv("S3_BUCKET_NAME"),
			os.Getenv("AWS_REGION"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Printf("Server starting on port :%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}