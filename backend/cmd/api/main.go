package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yourname/ory-fullstack/internal/auth"
	v1 "github.com/yourname/ory-fullstack/internal/handlers/api/v1"
	"github.com/yourname/ory-fullstack/pkg/config"
	"github.com/yourname/ory-fullstack/pkg/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewPostgresDB(cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Ory clients
	kratosClient := auth.NewKratosClient(cfg.KratosURL, cfg.KratosAdminURL)
	hydraClient := auth.NewHydraClient(cfg.HydraURL, cfg.HydraAdminURL)
	oathkeeperClient := auth.NewOathkeeperClient(cfg.OathkeeperURL)

	// Create Gin router
	r := gin.Default()

	// Initialize handlers
	authHandler := v1.NewAuthHandler(kratosClient, hydraClient, oathkeeperClient)

	// API routes
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
			auth.GET("/session", authHandler.GetSession)
		}

		user := api.Group("/user")
		{
			user.GET("/profile", authHandler.Authenticate, v1.GetProfile)
			user.PUT("/password", authHandler.Authenticate, v1.ChangePassword)
		}
	}

	// Start server
	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}