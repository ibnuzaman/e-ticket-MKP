package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"e-ticketing/handlers"
	"e-ticketing/middleware"
	"e-ticketing/models"
	"e-ticketing/seeders"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Terminal{})

	if err := seeders.SeedUsers(db); err != nil {
		log.Fatal("Failed to seed users:", err)
	}

	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	authHandler := handlers.NewAuthHandler(db)
	terminalHandler := handlers.NewTerminalHandler(db)

	e.POST("/login", authHandler.Login)
	api := e.Group("/api")
	api.Use(middleware.Auth(os.Getenv("JWT_SECRET")))
	api.POST("/terminals", terminalHandler.CreateTerminal)

	e.Logger.Fatal(e.Start(":8080"))
}
