package main

import (
	"fmt"
	"log"
	"missing-persons-backend/db"
)

func main (){


	database,err := db.Connect()


	if err !=nil {
		log.Fatalf("Database connection failed : %v", err)
	}

	defer database.Close ()

	fmt.Println ("Go server ready")
}