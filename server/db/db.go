package db

import (
    "fmt"
    "database/sql"
    "os"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
    if err := godotenv.Load(); err != nil {
        return nil, fmt.Errorf("error loading .env file: %w", err)
    }

    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_DATABASE")
    dbPort := os.Getenv("DB_PORT")

    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify",
        dbUser,
        dbPassword,
        dbHost,
        dbPort,
        dbName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error connecting to the database: %w", err)
    }

    fmt.Println("Successfully connected to the database")
    return db, nil
}