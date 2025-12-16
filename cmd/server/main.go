package main

import (
	"log"
	"os"

	"user-api/config"
	"user-api/internal/handler"
	"user-api/internal/logger"
	"user-api/internal/middleware"
	"user-api/internal/repository"
	"user-api/internal/routes"
	"user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file (development only)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	// Validate required env vars
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	app := fiber.New()

	// Custom logger (rename to avoid conflict)
	appLogger := logger.New()

	db := config.ConnectDB()

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc)

	app.Use(middleware.Logger(appLogger))

	routes.Setup(app, userHandler)

	log.Println("Server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
