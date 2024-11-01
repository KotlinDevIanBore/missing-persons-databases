package main

import (
	"fmt"
	"log"
	"missing-persons-backend/db"
	"missing-persons-backend/internal/handler"
	"missing-persons-backend/internal/service"
	"missing-persons-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

func main (){


	database,err := db.Connect()


	if err !=nil {
		log.Fatalf("Database connection failed : %v", err)
	}

	defer database.Close ()

	fmt.Println ("Go server ready")



	personRepo := &repository.PersonRepository{
		DB: database,
	}
	
	personService := &service.PersonService{

		Repo:personRepo,

	}

	personHandler := &handler.PersonHandler{

		Service :personService,
	}

	router :=  gin.Default ()

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


	router.GET ("api/persons", personHandler.GetMissingPersons)
	router.POST("api/persons",personHandler.CreateMissingPersons)

	port := ":8081"
	fmt.Printf("Server starting on port %s\n",port)

	if err := router.Run(port);err != nil {
		log.Fatalf( "failed to start server : %v", err)
	}


}

